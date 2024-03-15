package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	listingpb "github.com/pandishpancheta/api-gateway-service/pkg/listings/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UpdateListingRequest struct {
	Name        string
	Description string
	Price       string
	TagNames    []string
}

type UpdateStatusRequest struct {
	Status string
}

func GetAllListings(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient) {
	res, err := c.GetListings(r.Context(), &emptypb.Empty{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)

	w.WriteHeader(http.StatusOK)
}

func GetListing(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Call the service
	res, err := c.GetListing(r.Context(), &listingpb.GetListingRequest{Id: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func CreateListing(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient, userID string) {
	createListingRequest := &listingpb.CreateListingRequest{}

	// Decode the request body into the CreateListingRequest
	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// cast the file to a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createListingRequest.Chunk = fileBytes
	createListingRequest.Name = r.FormValue("name")
	createListingRequest.Price = r.FormValue("price")
	createListingRequest.UserId = userID

	// TagNames is a repeated field, so we need to loop through the form values and add them to the request
	for _, tag := range r.Form["tags"] {
		createListingRequest.TagNames = append(createListingRequest.TagNames, tag)
	}

	// Call the service
	res, err := c.CreateListing(r.Context(), createListingRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateListing(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient, userID string) {
	vars := mux.Vars(r)
	id := vars["id"]

	var body UpdateListingRequest

	// Decode the request body into the UpdateListingRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service
	res, err := c.UpdateListing(r.Context(), &listingpb.UpdateListingRequest{
		Id:          id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		TagNames:    body.TagNames,
		UserId:      userID,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateListingStatus(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient) {
	vars := mux.Vars(r)
	id := vars["id"]

	var body UpdateStatusRequest

	// Decode the request body into the UpdateStatusRequest
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service
	res, err := c.UpdateListingStatus(r.Context(), &listingpb.UpdateListingStatusRequest{
		Id:     id,
		Status: body.Status,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func DeleteListing(w http.ResponseWriter, r *http.Request, c listingpb.ListingsServiceClient, userID string) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Call the service
	res, err := c.DeleteListing(r.Context(), &listingpb.DeleteListingRequest{Id: id, UserId: userID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response as JSON to the writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
