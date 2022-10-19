//go:build wireinject
// +build wireinject

package server

import (
	"golayout/internal/service"
	"golayout/pkg/config"

	"github.com/google/wire"
)

func Init() *Server {
	wire.Build(
		configConstruct,
		wire.Value(service.Setups),
		NewServer,
	)

	return &Server{}
}

func configConstruct() *config.Config {
	cfg, err := config.New()
	if err != nil {
		return nil
	}

	return cfg
}
