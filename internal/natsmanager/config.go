package natsmanager

import "github.com/destinyarena/registration/internal/utils"

type (
	// Config holds configuration information for natsmanager
	Config struct {
		URL string `env:"NATS_URL,required"`
	}
)

func newEnvConfig() (*Config, error) {
	cfg := new(Config)
	if err := utils.EnvLoader(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
