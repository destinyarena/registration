package profilestore

import "github.com/sirupsen/logrus"

type (
	store struct {
		Config *Config
		Logger *logrus.Logger
	}

	Store interface {
		InsertUser(*User) (bool, error)
		GetUsersByIP(string) ([]*User, error)
	}
)

func New(logger *logrus.Logger, config *Config) (Store, error) {
	s := &store{
		Config: config,
		Logger: logger,
	}

	return s, nil
}

func NewDefault(logger *logrus.Logger) (Store, error) {
	config, err := newEnvConfig()
	if err != nil {
		return nil, err
	}

	return New(logger, config)
}
