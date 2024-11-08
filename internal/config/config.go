package config

import "github.com/amirmahdi-for/AemondTargaryen/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
