package routes

import (
	"encoding/json"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"net/http"
)

type LoginRequest struct {
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	var loginRequest LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.Login(r.Context(), &authpb.LoginRequest{Email: loginRequest.Email, Password: loginRequest.Password})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request, c authpb.AuthServiceClient) {
	var registerRequest RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.Register(r.Context(), &authpb.RegisterRequest{Username: registerRequest.Username, Email: registerRequest.Email, Password: registerRequest.Password})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
