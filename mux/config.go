package mux

import (
	"flag"
	"io/ioutil"

	"github.com/jasonrichardsmith/sentry/healthz"
	"github.com/jasonrichardsmith/sentry/tags"
	"github.com/jasonrichardsmith/sentry/limits"
	"github.com/jasonrichardsmith/sentry/sentry"
	yaml "gopkg.in/yaml.v2"
)

var (
	configfile string
)

func init() {
	flag.StringVar(&configfile, "sentry-config", "config.yaml", "Location of sentry config file.")
}

type Config struct {
	Limits  limits.Config  `yaml:"limits"`
	Healthz healthz.Config `yaml:"healthz"`
	Tags  tags.Config  `yaml:"tags"`
}

func New() *Config {
	l := limits.Config{}
	h := healthz.Config{}
	i := tags.Config{}
	return &Config{
		Limits:  l,
		Healthz: h,
		Tags:  i,
	}
}

func (c *Config) LoadFromFile() error {
	if !flag.Parsed() {
		flag.Parse()
	}
	configbytes, err := ioutil.ReadFile(configfile)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(configbytes, &c)
}

func (c *Config) LoadSentry() (sentry.Sentry, error) {
	return NewFromConfig(*c)
}
