package auth

import (
	"encoding/json"
	"net/http"
	_ "net/http"

	"github.com/gorilla/mux"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth/routes"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"

	_ "github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router, c *config.Config) *ServiceClient {
	authClient, err := InitAuthServiceClient(c)
	if err != nil {
		panic(err)
	}

	usersClient, err := InitUserServiceClient(c)
	if err != nil {
		panic(err)
	}

	svc := &ServiceClient{
		AuthClient:  authClient,
		UsersClient: usersClient,
	}

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", svc.Login).Methods("POST")
	auth.HandleFunc("/register", svc.Register).Methods("POST")
	auth.HandleFunc("/validate", svc.ValidateToken).Methods("POST")

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/{id}", svc.GetUser).Methods("GET")
	users.HandleFunc("/", svc.GetCurrentUser).Methods("GET")
	users.HandleFunc("/users/{id}", svc.DeleteCurrentUser).Methods("DELETE")

	return svc
}

func (svc *ServiceClient) Register(writer http.ResponseWriter, request *http.Request) {
	routes.Register(writer, request, svc.AuthClient)
}

func (svc *ServiceClient) Login(writer http.ResponseWriter, request *http.Request) {
	routes.Login(writer, request, svc.AuthClient)
}

func (svc *ServiceClient) ValidateToken(writer http.ResponseWriter, request *http.Request) {
	ValidateToken(writer, request, svc.AuthClient)
}

func (svc *ServiceClient) GetUser(writer http.ResponseWriter, request *http.Request) {
	routes.GetUser(writer, request, svc.UsersClient)
}

func (svc *ServiceClient) GetCurrentUser(writer http.ResponseWriter, request *http.Request) {
	routes.GetCurrentUser(writer, request, svc.UsersClient)
}

func (svc *ServiceClient) DeleteCurrentUser(writer http.ResponseWriter, request *http.Request) {
	routes.DeleteCurrentUser(writer, request, svc.UsersClient)
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
