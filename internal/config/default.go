package config

import "github.com/amirmahdi-for/AemondTargaryen/pkg/logger"

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
	}
}
