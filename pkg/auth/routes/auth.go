package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	authpb "github.com/pandishpancheta/api-gateway-service/pkg/auth/pb"
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

type GetUserRequest struct {
	Id string
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
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		log.Println(res)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request, c authpb.UserServiceClient) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := c.GetUser(r.Context(), &authpb.GetUserRequest{Id: id})
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

func GetUsers(w http.ResponseWriter, r *http.Request, c authpb.UserServiceClient) {
	res, err := c.GetUsers(r.Context(), &authpb.GetUsersRequest{})
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

func GetCurrentUser(w http.ResponseWriter, r *http.Request, c authpb.UserServiceClient) {
	token := r.Header.Get("Authorization")
	token = token[7:]

	res, err := c.GetCurrentUser(r.Context(), &authpb.GetCurrentUserRequest{Jwt: token})
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

func GetUserByUsername(w http.ResponseWriter, r *http.Request, c authpb.UserServiceClient) {
	vars := mux.Vars(r)
	username := vars["username"]

	res, err := c.GetUserByUsername(r.Context(), &authpb.GetUserByUsernameRequest{Username: username})
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

func DeleteCurrentUser(w http.ResponseWriter, r *http.Request, c authpb.UserServiceClient) {
	token := r.Header.Get("Authorization")
	token = token[7:]

	res, err := c.DeleteCurrentUser(r.Context(), &authpb.DeleteCurrentUserRequest{Jwt: token})
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
