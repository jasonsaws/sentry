package example

import (
	"github.com/jasonrichardsmith/sentry/sentry"
)

type Config struct {
	sentry.Config `yaml:"-,inline"`
}

func (c *Config) LoadSentry() (sentry.Sentry, error) {
	return ExampleSentry{}, nil
}