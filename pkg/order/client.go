package order

import (
	"fmt"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	orderpb "github.com/pandishpancheta/api-gateway-service/pkg/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client     orderpb.OrderServiceClient
	AuthClient authpb.AuthServiceClient
}

func InitServiceClient(cfg *config.Config) (orderpb.OrderServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.OrderServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to order service:", err)
		return nil, err
	}

	return orderpb.NewOrderServiceClient(c), nil
}
