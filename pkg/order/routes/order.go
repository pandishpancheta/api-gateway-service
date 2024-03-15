package routes

import (
	"encoding/json"
	orderpb "github.com/pandishpancheta/api-gateway-service/pkg/order/pb"
	"net/http"
)

type NewOrderRequest struct {
	ListingId string
	Status    string
	TokenUri  string
}

type GetOrdersByIDRequest struct {
	Id string
}

type UpdateStatusRequest struct {
	Id     string
	Status string
}

func CreateOrder(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient) {
	var newOrderRequest NewOrderRequest

	userId := r.Header.Get("user_id")

	if err := json.NewDecoder(r.Body).Decode(&newOrderRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.CreateOrder(r.Context(), &orderpb.NewOrderRequest{UserId: userId, ListingId: newOrderRequest.ListingId, Status: newOrderRequest.Status, TokenUri: newOrderRequest.TokenUri})
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

func GetOrdersByUser(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient) {
	userId := r.Header.Get("user_id")

	res, err := c.GetOrdersByUser(r.Context(), &orderpb.GetOrdersByUserRequest{UserId: userId})
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

func GetOrdersByID(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient) {
	var getOrdersByIDRequest GetOrdersByIDRequest

	if err := json.NewDecoder(r.Body).Decode(&getOrdersByIDRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.GetOrderByID(r.Context(), &orderpb.GetOrderByIDRequest{Id: getOrdersByIDRequest.Id})
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

func UpdateStatus(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient) {
	var updateStatusRequest UpdateStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&updateStatusRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.UpdateStatus(r.Context(), &orderpb.UpdateStatusRequest{Id: updateStatusRequest.Id, Status: updateStatusRequest.Status})
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