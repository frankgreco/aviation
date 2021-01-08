package grpc

import (
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/frankgreco/aviation/internal/run"
)

type Options struct {
	Server     *grpc.Server
	Network    string
	Address    string
	Reflection bool
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

	log.Printf("grpc server listening on %s\n", s.listener.Addr().String())

	return s.options.Server.Serve(s.listener)
}

func (s *grpcServer) Close(err error) error {
	log.Println("closing grpc server")
	defer s.options.Server.GracefulStop()
	return s.listener.Close()
}
