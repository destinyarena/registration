package profilestore

import (
	"context"

	pb "github.com/destinyarena/registration/pkg/profiles"
)

func (s *store) InsertUser(u *User) (bool, error) {
	client, conn, err := s.connect()
	if err != nil {
		return false, err
	}

	defer conn.Close()

	_, err = client.CreateProfile(context.Background(), &pb.ProfileRequest{
		Discord: u.Discord,
		Bungie:  u.Bungie,
		Faceit:  u.Faceit,
		Iphash:  u.IPHash,
	})

	if err != nil {
		return true, err
	}

	return false, nil
}
