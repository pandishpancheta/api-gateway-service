package routes

import (
	"encoding/json"
	"net/http"

	orderpb "github.com/pandishpancheta/api-gateway-service/pkg/order/pb"
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

func CreateOrder(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient, userId string) {
	var newOrderRequest NewOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&newOrderRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := c.CreateOrder(r.Context(), &orderpb.NewOrderRequest{UserId: userId, ListingId: newOrderRequest.ListingId})
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

func GetOrdersByUser(w http.ResponseWriter, r *http.Request, c orderpb.OrderServiceClient, userId string) {

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
