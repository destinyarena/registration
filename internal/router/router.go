package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// New returns a new echo router
func New(logger *logrus.Logger, config *Config) *echo.Echo {
	r := echo.New()

	r.Use(middleware.Recover())
	r.Use(middleware.Logger())

	return r
}

// NewDefault loads the config from env and returns it with the echo router
func NewDefault(logger *logrus.Logger) (*echo.Echo, *Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, nil, err
	}

	r := New(logger, cfg)

	return r, cfg, nil
}
