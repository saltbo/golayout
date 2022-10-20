package config

import (
	"golayout/pkg/database/mysql"

	"github.com/spf13/viper"
)

type Config struct {
	DB *mysql.Config
}

func New() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
