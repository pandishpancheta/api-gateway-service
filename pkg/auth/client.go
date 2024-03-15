package auth

import (
	"fmt"

	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	AuthClient  authpb.AuthServiceClient
	UsersClient authpb.UserServiceClient
}

func InitAuthServiceClient(cfg *config.Config) (authpb.AuthServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.AuthServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to auth service:", err)
		return nil, err
	}

	return authpb.NewAuthServiceClient(c), nil
}

func InitUserServiceClient(cfg *config.Config) (authpb.UserServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.AuthServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to user service:", err)
		return nil, err
	}

	return authpb.NewUserServiceClient(c), nil
}
