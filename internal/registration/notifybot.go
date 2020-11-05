package registration

import "github.com/destinyarena/registration/internal/natsmanager"

func (h *handler) notifyBot(u *user) {
	payload := &natsmanager.RegistrationChannel{
		ID: u.Discord,
	}

	h.Logger.Info("Sending user payload to NATS")
	h.NATSChannels.Registration <- payload
}
