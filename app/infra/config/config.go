package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ubaidillahhf/go-clarch/app/infra/exception"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
