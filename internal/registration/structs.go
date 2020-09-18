package registration

type (
	// Payload is the jwt json payload
	Payload struct {
		Discord string `validate:"required"`
		Bungie  string `validate:"required"`
		Faceit  string `validate:"required"`
	}

	user struct {
		Discord string
		Bungie  string
		Faceit  string
		IPHash  string
	}
)
