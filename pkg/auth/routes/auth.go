package routes

import (
	"encoding/json"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
	"io/ioutil"
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

type ValidateTokenRequest struct {
	Token string
}

func Login(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var loginRequest LoginRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &loginRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.Login(r.Context(), &pb.LoginRequest{Email: loginRequest.Email, Password: loginRequest.Password})
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

func Register(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var registerRequest RegisterRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &registerRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.Register(r.Context(), &pb.RegisterRequest{Username: registerRequest.Username, Email: registerRequest.Email, Password: registerRequest.Password})
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

func ValidateToken(w http.ResponseWriter, r *http.Request, c pb.AuthServiceClient) {
	var validateTokenRequest ValidateTokenRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &validateTokenRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.ValidateToken(r.Context(), &pb.ValidateTokenRequest{Jwt: validateTokenRequest.Token})
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
