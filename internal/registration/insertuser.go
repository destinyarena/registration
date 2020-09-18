package registration

import (
	"github.com/destinyarena/registration/internal/profilestore"
)

func (h *handler) insertUser(u *user) (bool, error) {
	dbuser := &profilestore.User{
		Discord: u.Discord,
		Bungie:  u.Bungie,
		Faceit:  u.Faceit,
		IPHash:  u.IPHash,
	}

	if _, err := h.ProfileStore.InsertUser(dbuser); err != nil {
		return true, nil
	}

	return false, nil
}
