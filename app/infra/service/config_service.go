package service

import (
	"github.com/ubaidillahhf/go-clarch/app/domain"
	"github.com/ubaidillahhf/go-clarch/app/infra/config"
)

type configProvider struct {
	config config.IConfig
}

func NewConfigProvider(cfg config.IConfig) domain.ConfigProvider {
	return &configProvider{
		config: cfg,
	}
}

func (c *configProvider) Get(key string) string {
	return c.config.Get(key)
}
