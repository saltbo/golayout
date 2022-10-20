package service

import "github.com/google/wire"

type Services struct {
	*Book
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewBook,
	wire.Struct(new(Services), "*"),
)
