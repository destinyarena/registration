package registration

func (h *handler) isBanned(iphash string) (bool, error) {
	users, err := h.ProfileStore.GetUsersByIP(iphash)
	if err != nil {
		return false, err
	}

	h.Logger.Infof("Found %d users with this ip hash: %s", len(users), iphash)
	for _, u := range users {
		if u.Banned {
			return true, nil
		}
	}

	return false, nil

}
