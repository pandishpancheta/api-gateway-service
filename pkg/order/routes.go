package order

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/order/routes"
)

func RegisterRouters(r *mux.Router, cfg *config.Config, authClient authpb.AuthServiceClient) *ServiceClient {
	client, err := InitServiceClient(cfg)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		Client:     client,
		AuthClient: authClient,
	}

	router := r.PathPrefix("/orders").Subrouter()
	router.HandleFunc("/orders", svc.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", svc.GetOrdersByUser).Methods("GET")
	router.HandleFunc("/orders/{id}", svc.GetOrder).Methods("GET")
	router.HandleFunc("/orders/{id}", svc.UpdateStatus).Methods("PUT")

	return svc
}

func (svc *ServiceClient) CreateOrder(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.CreateOrder(writer, request, svc.Client, userId)
}

func (svc *ServiceClient) GetOrdersByUser(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.GetOrdersByUser(writer, request, svc.Client, userId)
}

func (svc *ServiceClient) GetOrder(writer http.ResponseWriter, request *http.Request) {
	routes.GetOrdersByID(writer, request, svc.Client)
}

func (svc *ServiceClient) UpdateStatus(writer http.ResponseWriter, request *http.Request) {
	routes.UpdateStatus(writer, request, svc.Client)
}
