package registration

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) endpoint(c echo.Context) error {
	payload := new(Payload)
	if err := c.Bind(payload); err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Validation TBD

	h.Logger.Info("Checking for bad IP's")
	badip, err := h.badIP(c.RealIP())
	if err != nil {
		h.Logger.Error(err)
		badip = false
	}

	if badip {
		err = errors.New("looks like you are trying to signup using a VPN, please contact an admin if this is a mistake")
		h.Logger.Error(err)
		return c.String(http.StatusForbidden, err.Error())
	}

	h.Logger.Info("Starting Alt Detection..")
	user, err := h.getUser(payload, c.RealIP())
	if err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	// IP detection
	banned, err := h.isBanned(user.IPHash)
	if err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if banned {
		err = errors.New("sorry looks like you are IP banned if you, please contact an admin if this is a mistake")
		h.Logger.Error(err)
		return c.String(http.StatusForbidden, err.Error())
	}

	// Alt detection

	alt, err := h.insertUser(user)
	if err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if alt {
		err = errors.New("sorry but an account with this information already exists, please contact and admin if this is a mistake")
		h.Logger.Error(err)
		return c.String(http.StatusForbidden, err.Error())
	}

	h.Logger.Info("Notifing discord bot")
	if err = h.discordHandler(user); err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "You have successfully registered")
}
