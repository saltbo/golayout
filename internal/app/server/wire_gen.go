// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"golayout/internal/service"
	"golayout/pkg/config"
)

// Injectors from wire.go:

func Init() *Server {
	config := configConstruct()
	v := _wireValue
	server := NewServer(config, v)
	return server
}

var (
	_wireValue = service.Setups
)

// wire.go:

func configConstruct() *config.Config {
	cfg, err := config.New()
	if err != nil {
		return nil
	}

	return cfg
}