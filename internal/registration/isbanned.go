package registration

func (h *handler) isBanned(iphash string) (bool, error) {
	users, err := h.ProfileStore.GetUsersByIP(iphash)
	if err != nil {
		return false, err
	}

	for _, u := range users {
		if u.Banned == true {
			return true, nil
		}
	}

	return false, nil

}
