package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	userspb "github.com/pandishpancheta/api-gateway-service/pkg/users/pb"
	"net/http"
)

type GetUserRequest struct {
	Id string
}

func GetUser(w http.ResponseWriter, r *http.Request, c userspb.UserServiceClient) {
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := c.GetUser(r.Context(), &userspb.GetUserRequest{Id: id})
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

func GetCurrentUser(w http.ResponseWriter, r *http.Request, c userspb.UserServiceClient) {
	token := r.Header.Get("Authorization")
	token = token[7:]

	res, err := c.GetCurrentUser(r.Context(), &userspb.GetCurrentUserRequest{Jwt: token})
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

func DeleteCurrentUser(w http.ResponseWriter, r *http.Request, c userspb.UserServiceClient) {
	token := r.Header.Get("Authorization")
	token = token[7:]

	res, err := c.DeleteCurrentUser(r.Context(), &userspb.DeleteCurrentUserRequest{Jwt: token})
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
