package auth

import (
	"fmt"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (pb.AuthServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.AuthServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to auth service:", err)
		return nil, err
	}

	return pb.NewAuthServiceClient(c), nil
}
