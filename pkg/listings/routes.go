package listings

import (
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
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
	svc := &ServiceClient{
		Client: client,
	}

	router := r.PathPrefix("/listings").Subrouter()
	router.HandleFunc("/listings", svc.GetListings).Methods("GET")
	router.HandleFunc("/listings", svc.CreateListing).Methods("POST")
	router.HandleFunc("/listings/{id}", svc.GetListing).Methods("GET")
	router.HandleFunc("/listings/{id}", svc.UpdateListing).Methods("PUT")
	router.HandleFunc("/listings/{id}", svc.DeleteListing).Methods("DELETE")

	return svc
}

func (svc *ServiceClient) DeleteListing(writer http.ResponseWriter, request *http.Request) {
	routes.DeleteListing(writer, request, svc.Client)
}

func (svc *ServiceClient) UpdateListing(writer http.ResponseWriter, request *http.Request) {
	routes.UpdateListing(writer, request, svc.Client)
}

func (svc *ServiceClient) GetListing(writer http.ResponseWriter, request *http.Request) {
	routes.GetListing(writer, request, svc.Client)
}

func (svc *ServiceClient) CreateListing(writer http.ResponseWriter, request *http.Request) {
	routes.CreateListing(writer, request, svc.Client)
}

func (svc *ServiceClient) GetListings(writer http.ResponseWriter, request *http.Request) {
	routes.GetAllListings(writer, request, svc.Client)
}
