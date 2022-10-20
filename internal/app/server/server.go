package server

import (
	"context"

	v1 "golayout/api/helloworld/v1"
	"golayout/internal/pkg/config"
	"golayout/internal/service"

	"google.golang.org/grpc"
)

type Server struct {
	cfg *config.Config

	grpcSrv *grpc.Server
}

func NewServer() (*Server, error) {
	return wireNewServer()
}

func newServer(cfg *config.Config, grpcSrv *grpc.Server, ss *service.Services) (*Server, error) {
	v1.RegisterGreeterServer(grpcSrv, ss.Book)

	return &Server{
		cfg:     cfg,
		grpcSrv: grpcSrv,
	}, nil
}

func (s *Server) Run(ctx context.Context) error {

	return nil
}
