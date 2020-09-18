package natsmanager

type (
	// RegistrationChannel is the channel used for regisrtation messaging
	RegistrationChannel struct {
		ID string `json:"id"`
	}

	// Channels holds pointers to all created channels
	Channels struct {
		Registration chan *RegistrationChannel
	}
)
