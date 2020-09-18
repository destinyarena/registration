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

	user, err := h.getUser(payload, c.RealIP())
	if err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusUnauthorized, err.Error())
	}

	alt, err := h.insertUser(user)
	if alt {
		err = errors.New("Sorry but an account with this information already exists of you have been banned, please contact and admin if this is a mistake.")
		h.Logger.Error(err)
		return c.String(http.StatusForbidden, err.Error())
	}

	if err != nil {
		h.Logger.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	h.notifyBot(user)

	return c.String(http.StatusOK, "You have successfully registered")
}
