package registration

import (
	"github.com/destinyarena/registration/internal/jwtmanager"
	"github.com/destinyarena/registration/internal/natsmanager"
	"github.com/destinyarena/registration/internal/profilestore"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	handler struct {
		Logger       *logrus.Logger
		ProfileStore profilestore.Store
		JWTManager   jwtmanager.JWTManager
		NATSChannels *natsmanager.Channels
	}

	Handler interface {
		Register(*echo.Group) error
	}
)

func (h *handler) Register(g *echo.Group) error {
	g.POST("", h.endpoint)
	return nil
}

func NewDefault(logger *logrus.Logger, jwtmanager jwtmanager.JWTManager, profileStore profilestore.Store, natsChannels *natsmanager.Channels) (Handler, error) {
	h := &handler{
		Logger:       logger,
		ProfileStore: profileStore,
		JWTManager:   jwtmanager,
		NATSChannels: natsChannels,
	}

	return h, nil
}
