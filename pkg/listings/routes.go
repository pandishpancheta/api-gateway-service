package listings

import (
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/listings/routes"
	"net/http"
	_ "net/http"
)

func RegisterRouters(r *mux.Router, c *config.Config) *ServiceClient {
	client, err := InitServiceClient(c)
	if err != nil {
		panic(err)
	}

	authClient, err := auth.InitServiceClient(c)

	svc := &ServiceClient{
		Client:     client,
		AuthClient: authClient,
	}

	router := r.PathPrefix("/listings").Subrouter()
	router.HandleFunc("/", svc.GetListings).Methods("GET")
	router.HandleFunc("/", svc.CreateListing).Methods("POST")
	router.HandleFunc("/{id}", svc.GetListing).Methods("GET")
	router.HandleFunc("/{id}", svc.UpdateListing).Methods("PUT")
	router.HandleFunc("/{id}", svc.DeleteListing).Methods("DELETE")

	return svc
}

func (svc *ServiceClient) DeleteListing(writer http.ResponseWriter, request *http.Request) {
	auth.ValidateToken(writer, request, svc.AuthClient)
	routes.DeleteListing(writer, request, svc.Client)
}

func (svc *ServiceClient) UpdateListing(writer http.ResponseWriter, request *http.Request) {
	auth.ValidateToken(writer, request, svc.AuthClient)
	routes.UpdateListing(writer, request, svc.Client)
}

func (svc *ServiceClient) GetListing(writer http.ResponseWriter, request *http.Request) {
	routes.GetListing(writer, request, svc.Client)
}

func (svc *ServiceClient) CreateListing(writer http.ResponseWriter, request *http.Request) {
	auth.ValidateToken(writer, request, svc.AuthClient)
	routes.CreateListing(writer, request, svc.Client)
}

func (svc *ServiceClient) GetListings(writer http.ResponseWriter, request *http.Request) {
	routes.GetAllListings(writer, request, svc.Client)
}
