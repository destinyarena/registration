package registration

import (
	"github.com/destinyarena/registration/internal/profilestore"
)

func (h *handler) insertUser(u *user) (bool, error) {
	banned, err := h.isBanned(u.IPHash)
	if err != nil {
		return true, err
	}

	if banned {
		return true, nil
	}

	dbuser := &profilestore.User{
		Discord: u.Discord,
		Bungie:  u.Bungie,
		Faceit:  u.Faceit,
		IPHash:  u.IPHash,
	}

	return h.ProfileStore.InsertUser(dbuser)
}
