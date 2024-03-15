package server

import (
	"github.com/caarlos0/env/v10"
	"github.com/pkg/errors"
)

func setupConfig(s *Server) error {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return errors.Wrap(err, "parsing config")
	}

	if err := c.Validate(); err != nil {
		return errors.Wrap(err, "failed to validate config")
	}

	s.Config = c
	return nil
}

type Config struct {
	Name              string `env:"NAME" envDefault:"codex"`
	Mongo             string `env:"MONGO" envDefault:"mongodb://localhost:27017"`
	Port              string `env:"PORT" envDefault:"9020"`
	Production        bool   `env:"PRODUCTION" envDefault:"false"`
	PlexToken         string `env:"PLEX_TOKEN"`
	PlexURL           string `env:"PLEX_URL"`
	MinionConcurrency int    `env:"MINION_CONCURRENCY" envDefault:"10"`
	MinionDebug       bool   `env:"MINION_DEBUG" envDefault:"false"`
	MinionBufferSize  int    `env:"MINION_BUFFER_SIZE" envDefault:"100"`
	MinionURI         string `env:"MINION_URI"`
	MinionDatabase    string `env:"MINION_DATABASE"`
	MinionCollection  string `env:"MINION_COLLECTION"`
}

func (c *Config) Validate() error {
	return nil
}
