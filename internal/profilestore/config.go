package profilestore

import "github.com/destinyarena/registration/internal/utils"

type (
	Config struct {
		Host string `env:"PROFILES_HOST,required"`
	}
)

func newEnvConfig() (*Config, error) {
	config := new(Config)

	if err := utils.EnvLoader(config); err != nil {
		return nil, err
	}

	return config, nil
}
