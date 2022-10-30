package registration

// This whole project is getting refactored soon so this will be fixed then.
const (
	roleID  = "738885109494841394"
	guildID = "650109209610027034"
	invite  = "https://www.faceit.com/en/inv/bHFVEo3"
)

func (h *handler) discordHandler(u *user) error {
	h.Logger.Infof("Discord: %s", u.Discord)

	// Get the guild
	err := h.DSession.GuildMemberRoleAdd(guildID, u.Discord, roleID) // Add the role to the user
	if err != nil {
		return err
	}

	_, err = h.DSession.ChannelMessageSend(u.Discord, "Welcome to Destiny Arena! Join our Faceit hub: "+invite) // Send a message to the user
	if err != nil {
		return err
	}

	return nil
}
