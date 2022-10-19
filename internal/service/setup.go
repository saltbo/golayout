package service

import (
	"context"
)

type SvcSetup interface {
	Setup(ctx context.Context) error
}

var Setups = []SvcSetup{
	NewBook(),
}
