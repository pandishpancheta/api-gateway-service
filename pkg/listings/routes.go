package listings

import (
	"net/http"
	_ "net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/listings/routes"
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

	router := r.PathPrefix("/listings").Subrouter()
	router.HandleFunc("/", svc.GetListings).Methods("GET")
	router.HandleFunc("/", svc.CreateListing).Methods("POST")
	router.HandleFunc("/{id}", svc.GetListing).Methods("GET")
	router.HandleFunc("/{id}", svc.UpdateListing).Methods("PUT")
	router.HandleFunc("/{id}", svc.DeleteListing).Methods("DELETE")

	return svc
}

func (svc *ServiceClient) DeleteListing(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.DeleteListing(writer, request, svc.Client, userId)
}

func (svc *ServiceClient) UpdateListing(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.UpdateListing(writer, request, svc.Client, userId)
}

func (svc *ServiceClient) GetListing(writer http.ResponseWriter, request *http.Request) {
	routes.GetListing(writer, request, svc.Client)
}

func (svc *ServiceClient) CreateListing(writer http.ResponseWriter, request *http.Request) {
	userId, err := auth.ValidateToken(writer, request, svc.AuthClient)
	if err != nil {
		return
	}
	routes.CreateListing(writer, request, svc.Client, userId)
}

func (svc *ServiceClient) GetListings(writer http.ResponseWriter, request *http.Request) {
	routes.GetAllListings(writer, request, svc.Client)
}
