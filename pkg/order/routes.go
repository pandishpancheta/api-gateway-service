package order

import (
	"github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/order/routes"
	"net/http"
)

func RegisterRouters(r *mux.Router, c *config.Config) *ServiceClient {
	client, err := InitServiceClient(c)
	if err != nil {
		panic(err)
	}

	authClient, err := auth.InitServiceClient(c)
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
	auth.ValidateToken(writer, request, svc.AuthClient)
	routes.CreateOrder(writer, request, svc.Client)
}

func (svc *ServiceClient) GetOrdersByUser(writer http.ResponseWriter, request *http.Request) {
	auth.ValidateToken(writer, request, svc.AuthClient)
	routes.GetOrdersByUser(writer, request, svc.Client)
}

func (svc *ServiceClient) GetOrder(writer http.ResponseWriter, request *http.Request) {
	routes.GetOrdersByID(writer, request, svc.Client)
}

func (svc *ServiceClient) UpdateStatus(writer http.ResponseWriter, request *http.Request) {
	routes.UpdateStatus(writer, request, svc.Client)
}
