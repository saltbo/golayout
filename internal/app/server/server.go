package server

import (
	"context"

	"golayout/internal/service"
	"golayout/pkg/config"
)

type Server struct {
	cfg *config.Config

	setups []service.SvcSetup
}

func NewServer(cfg *config.Config, setups []service.SvcSetup) *Server {
	return &Server{
		cfg:    cfg,
		setups: setups,
	}
}

func (s *Server) Run(ctx context.Context) error {
	for _, setup := range s.setups {
		setup.Setup(ctx)
	}

	return nil
}
