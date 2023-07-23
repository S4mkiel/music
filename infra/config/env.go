package config

import (
	env "github.com/Netflix/go-env"
	"github.com/S4mkiel/music/infra/db"
	"github.com/S4mkiel/music/infra/http"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("config",
	fx.Provide(NewConfig),
	fx.Provide(func(cfg Config) http.Config { return cfg.Http }),
	fx.Provide(func(cfg Config) db.Config { return cfg.Db }),
)

type Config struct {
	Extras *env.EnvSet
	Http   http.Config
	Db     db.Config
}

func NewConfig(logger *zap.SugaredLogger) Config {
	var cfg Config = Config{}
	err := cfg.loadConfig()
	if err != nil {
		logger.Error(err)
	}

	return cfg
}

func (c *Config) loadConfig() error {
	es, err := env.UnmarshalFromEnviron(c)
	if err != nil {
		return err
	}
	c.Extras = &es

	return nil
}
