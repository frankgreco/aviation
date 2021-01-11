package grpc

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/frankgreco/aviation/internal/log"
	"github.com/frankgreco/aviation/internal/run"
)

type Options struct {
	Server     *grpc.Server
	Network    string
	Address    string
	Reflection bool
	Logger     log.Logger
}

type grpcServer struct {
	options  *Options
	listener net.Listener
	err      error
}

func Prepare(options *Options) run.Runnable {
	s := &grpcServer{
		options: options,
	}

	if s.options.Server == nil {
		s.err = errors.New("server must be defined")
	}

	if s.options.Reflection {
		reflection.Register(s.options.Server)
	}

	return s
}

func (s *grpcServer) Run() error {
	if s.err != nil {
		return s.err
	}

	s.listener, s.err = net.Listen(s.options.Network, s.options.Address)
	if s.err != nil {
		return s.err
	}

	s.options.Logger.Info(fmt.Sprintf("grpc server listening on %s", s.listener.Addr().String()))
	return s.options.Server.Serve(s.listener)
}

func (s *grpcServer) Close(error) error {
	if s.options.Server != nil {
		s.options.Server.GracefulStop()
	}
	if s.listener != nil {
		// https://golang.org/src/net/error_test.go#L501
		if err := s.listener.Close(); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			s.options.Logger.Error(fmt.Sprintf("error closing gRPC listener: %s", err.Error()))
			return err
		}
	}
	s.options.Logger.Info("closed grpc server")
	return nil
}
