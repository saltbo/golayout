//go:build wireinject
// +build wireinject

package server

import (
	"golayout/internal/dao"
	"golayout/internal/pkg/config"
	"golayout/internal/service"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

func wireNewServer(option ...grpc.ServerOption) (*Server, error) {
	wire.Build(
		config.New,
		dao.ProviderSet,
		service.ProviderSet,
		grpc.NewServer,
		newServer,
	)

	return &Server{}, nil
}
