package profilestore

import (
	pb "github.com/destinyarena/registration/pkg/profiles"
	"google.golang.org/grpc"
)

func (s *store) connect() (pb.ProfilesClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(s.Config.Host, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewProfilesClient(conn)

	return client, conn, nil
}
