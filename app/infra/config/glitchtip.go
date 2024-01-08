package config

import (
	"github.com/getsentry/sentry-go"
)

func GlitchtipInit(configuration IConfig) {
	sentry.Init(sentry.ClientOptions{
		Dsn:              configuration.Get("GLITCHTIP_DSN"),
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
}
