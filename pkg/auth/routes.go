package auth

import (
	"github.com/gorilla/mux"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth/routes"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"net/http"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router, c *config.Config) *ServiceClient {
	client, err := InitServiceClient(c)
	if err != nil {
		panic(err)
	}
	svc := &ServiceClient{
		Client: client,
	}

	router := r.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login", svc.Login).Methods("POST")
	router.HandleFunc("/register", svc.Register).Methods("POST")
	router.HandleFunc("/validate", svc.ValidateToken).Methods("POST")

	return svc
}

func (svc *ServiceClient) Register(writer http.ResponseWriter, request *http.Request) {
	routes.Register(writer, request, svc.Client)
}

func (svc *ServiceClient) Login(writer http.ResponseWriter, request *http.Request) {
	routes.Login(writer, request, svc.Client)
}

func (svc *ServiceClient) ValidateToken(writer http.ResponseWriter, request *http.Request) {
	routes.ValidateToken(writer, request, svc.Client)
}
