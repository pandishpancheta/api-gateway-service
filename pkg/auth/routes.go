package auth

import (
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http"

	_ "github.com/gorilla/mux"
)

func RegisterRouters(r *mux.Router) {
	r.HandleFunc("/login", Login).Methods("POST")
	r.HandleFunc("/register", Register).Methods("POST")
}

func Register(writer http.ResponseWriter, request *http.Request) {

}

func Login(writer http.ResponseWriter, request *http.Request) {

}
