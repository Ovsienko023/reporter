package infrastructure

import (
	"embed"
	"go.uber.org/zap"
)

type Infrastructure struct {
	embedFS embed.FS
	logger  *zap.Logger
}

type Config struct {
	Logger *zap.Logger
}

func New(config *Config) (*Infrastructure, error) {
	return &Infrastructure{
		logger: config.Logger,
	}, nil
}

func (i *Infrastructure) GetLogger() *zap.Logger {
	return i.logger
}
