package dao

import (
	"golayout/pkg/database/mysql"

	"github.com/google/wire"
)

// Dao .
type Dao struct {
	db *mysql.DB
}

// NewDao .
func NewDao(c *mysql.Config) (*Dao, func(), error) {
	cleanup := func() {
		// log.NewHelper(logger).Info("closing the data resources")
	}

	db, err := mysql.NewDB(c)
	if err != nil {
		return nil, nil, err
	}

	return &Dao{db}, cleanup, nil
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewDao,
	NewBook,
)
