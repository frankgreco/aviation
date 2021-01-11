package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"

	"github.com/frankgreco/aviation/internal/log"
)

type serverParams struct {
	secureServer   *http.Server
	insecureServer *http.Server
	err            []string
	options        *Options
}

// Options contains configuration for http(s) server(s)
type Options struct {
	Name                   string
	InsecureAddr           string
	SecureAddr             string
	InsecurePort           int
	SecurePort             int
	TLSKey, TLSCert, TLSCa string
	Handler                http.Handler
	Logger                 log.Logger
}

// Prepare will construct net.Listerer instantiations for the requested
// server(s). If an error is encounted, it will be hidden within the returned
// type for evaluation by that type's methods.
func Prepare(opts *Options) *serverParams {
	f, err := os.Open(opts.TLSCa)
	if err != nil && len(opts.TLSCa) > 0 {
		return &serverParams{err: []string{err.Error()}}
	} else if err == nil && len(opts.TLSCa) > 0 {
		defer func() {
			if err := f.Close(); err != nil {
				opts.Logger.Error(fmt.Sprintf("error closing %s: %s", opts.Name, err))
			}
		}()
	}
	return prepare(opts, f)
}

func prepare(opts *Options, r io.Reader) *serverParams {
	params := &serverParams{options: opts}
	insecureAddr, secureAddr := fmt.Sprintf("%s:%d", opts.InsecureAddr, opts.InsecurePort), fmt.Sprintf("%s:%d", opts.SecureAddr, opts.SecurePort)

	if opts.SecurePort > 0 && len(opts.TLSCa) > 0 { // client has requested an HTTPS server with mutual TLS
		tlsConfig, err := getTLSConfigFromReader(r)
		if err != nil {
			params.err = append(params.err, err.Error())
		} else {
			params.secureServer = &http.Server{Addr: secureAddr, TLSConfig: tlsConfig, Handler: opts.Handler}
		}
	} else if opts.SecurePort > 0 && len(opts.TLSCa) < 1 { // client has requested an HTTPS server with one-way TLS
		params.secureServer = &http.Server{Addr: secureAddr, Handler: opts.Handler}
	}
	if opts.InsecurePort > 0 { // client has requested an HTTP server
		params.insecureServer = &http.Server{Addr: insecureAddr, Handler: opts.Handler}
	}
	return params
}

// Run will start http(s) server(s) according to the reciever's configuration.
// If an error occurred durring the receiver's construction, that error will be
// returned immedietaly. Otherwise, run will return the non-nill error from the
// the first server that terminates.
func (params *serverParams) Run() error {
	if params.err != nil && len(params.err) > 0 {
		return errors.New(strings.Join(params.err, ", "))
	}

	var g errgroup.Group
	if params.secureServer != nil {
		params.options.Logger.Info(fmt.Sprintf("starting %s %s server on %s:%d", params.options.Name, "HTTPS", params.options.SecureAddr, params.options.SecurePort))
		g.Go(func() error {
			return params.secureServer.ListenAndServeTLS(params.options.TLSCert, params.options.TLSKey)
		})
	}
	if params.insecureServer != nil {
		params.options.Logger.Info(fmt.Sprintf("starting %s %s server on %s:%d", params.options.Name, "HTTP", params.options.InsecureAddr, params.options.InsecurePort))
		g.Go(func() error {
			return params.insecureServer.ListenAndServe()
		})
	}
	err := g.Wait()
	if err != nil && err != http.ErrServerClosed {
		params.options.Logger.Error(err.Error())
		return err
	}
	return nil
}

func (params *serverParams) IsEnabled() bool {
	if params.err != nil && len(params.err) > 0 {
		return false
	}
	return params.options.InsecurePort > 0 || params.options.SecurePort > 0
}

// Name returns the name of the server
func (params *serverParams) Name() string {
	return params.options.Name
}

// Close will gracefully terminate http(s) server(s) that were bootstrapped
// according to the reciever's configuration. More details here:
// https://github.com/golang/go/blob/master/src/net/http/server.go#L2545-L2561
func (params *serverParams) Close(error) error {
	if err := closeServer(params.options.Logger, params.secureServer, params.options.Name, "HTTPS"); err != nil {
		params.err = append(params.err, err.Error())
	}
	if err := closeServer(params.options.Logger, params.insecureServer, params.options.Name, "HTTP"); err != nil {
		params.err = append(params.err, err.Error())
	}
	if params.err == nil || len(params.err) == 0 {
		return nil
	}
	return errors.New(strings.Join(params.err, ", "))
}

func closeServer(log log.Logger, svr *http.Server, name, scheme string) error {
	if svr == nil {
		return nil
	}
	if err := svr.Shutdown(context.Background()); err != nil {
		log.Error(fmt.Sprintf("error gracefully closing %s %s server: %v", name, scheme, err))
		return err
	}
	log.Info(fmt.Sprintf("gracefully closed %s %s server", name, scheme))
	return nil
}

func getTLSConfigFromReader(r io.Reader) (*tls.Config, error) {
	if r == nil {
		return nil, errors.New("reader is nil")
	}
	caCert, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("could not append cert to pool")
	}
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}
