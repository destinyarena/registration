package profilestore

import (
	"context"
	"io"

	pb "github.com/destinyarena/registration/pkg/profiles"
)

func (s *store) GetUsersByIP(iphash string) ([]*User, error) {
	client, conn, err := s.connect()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	stream, err := client.GetProfilesByIP(context.Background(), &pb.IPRequest{
		Iphash: iphash,
	})

	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)

	for {
		u, err := stream.Recv()
		if err == io.EOF {
			s.Logger.Infof("End of Profiles stream for: %s", iphash)
			break
		}

		if err != nil {
			s.Logger.Errorf("Error: %s", err.Error())
			return nil, err
		}

		users = append(users, &User{
			Discord: u.Discord,
			Bungie:  u.Bungie,
			Faceit:  u.Faceit,
			Banned:  u.Banned,
			IPHash:  u.Iphash,
		})
	}

	return users, nil
}
