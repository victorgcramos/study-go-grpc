package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/victorgcramos/zi/pkg/config"
	"github.com/victorgcramos/zi/pkg/user/pb"
	"google.golang.org/grpc"
	insec "google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	client pb.UserServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	//	use insecure dial for testing purposes.
	d := grpc.WithTransportCredentials(insec.NewCredentials())

	cc, err := grpc.Dial(c.UserSvcUrl, d)

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{pb.NewUserServiceClient(cc)}
}

func (sc *ServiceClient) Register(ctx context.Context, nu NewUser) ([]byte, error) {
	res, err := sc.client.Register(ctx, &pb.RegisterRequest{
		Username: nu.Username,
		Password: nu.Password,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return json.Marshal(res)
}

func (sc *ServiceClient) Version(ctx context.Context) ([]byte, error) {
	res, err := sc.client.Version(ctx, &pb.VersionRequest{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return json.Marshal(res)
}
