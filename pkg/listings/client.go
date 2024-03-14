package listings

import (
	"fmt"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/listings/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.ListingsServiceClient
}

func InitServiceClient(cfg *config.Config) (pb.ListingsServiceClient, error) {
	creds := insecure.NewCredentials()
	c, err := grpc.Dial(cfg.ListingsServiceAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		fmt.Println("Error connecting to auth service:", err)
		return nil, err
	}

	return pb.NewListingsServiceClient(c), nil
}
