package users

import (
	"fmt"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	userspb "github.com/pandishpancheta/api-gateway-service/pkg/users/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client     userspb.UserServiceClient
	AuthClient authpb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (userspb.UserServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.UserServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to user service:", err)
		return nil, err
	}

	return userspb.NewUserServiceClient(c), nil
}
