package auth

import (
	"encoding/json"
	"github.com/gorilla/mux"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
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
	ValidateToken(writer, request, svc.Client)
}

func ValidateToken(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token = token[7:]

	res, err := c.ValidateToken(r.Context(), &authpb.ValidateTokenRequest{Jwt: token})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	r.Header.Set("user_id", res.UserId)

	w.WriteHeader(http.StatusOK)
}
