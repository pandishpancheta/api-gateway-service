package users

import (
	"github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/users/routes"
	"net/http"
)

func RegisterRoutes(r *mux.Router, c *config.Config) *ServiceClient {
	client, err := InitServiceClient(c)
	if err != nil {
		panic(err)
	}

	authClient, err := auth.InitServiceClient(c)

	svc := &ServiceClient{
		Client:     client,
		AuthClient: authClient,
	}

	router := r.PathPrefix("/users").Subrouter()
	router.HandleFunc("/{id}", svc.GetUser).Methods("GET")
	router.HandleFunc("/", svc.GetCurrentUser).Methods("GET")
	router.HandleFunc("/users/{id}", svc.DeleteCurrentUser).Methods("DELETE")

	return svc
}

func (svc *ServiceClient) GetUser(writer http.ResponseWriter, request *http.Request) {
	routes.GetUser(writer, request, svc.Client)
}

func (svc *ServiceClient) GetCurrentUser(writer http.ResponseWriter, request *http.Request) {
	routes.GetCurrentUser(writer, request, svc.Client)
}

func (svc *ServiceClient) DeleteCurrentUser(writer http.ResponseWriter, request *http.Request) {
	routes.DeleteCurrentUser(writer, request, svc.Client)
}
