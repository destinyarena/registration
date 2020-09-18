package natsmanager

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

const (
	stanClient = "arena-registration"
)

func NewDefault(log *logrus.Logger) (*Channels, error) {
	config, err := newEnvConfig()
	if err != nil {
		return nil, err
	}

	return New(config, log)
}

func New(config *Config, log *logrus.Logger) (*Channels, error) {
	nc, err := nats.Connect(config.URL)
	if err != nil {
		return nil, err
	}

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	regChan := make(chan *RegistrationChannel)

	ec.BindSendChan("registration", regChan)

	return &Channels{
		Registration: regChan,
	}, nil
}
