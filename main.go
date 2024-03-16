package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pandishpancheta/api-gateway-service/pkg/auth"
	"github.com/pandishpancheta/api-gateway-service/pkg/config"
	"github.com/pandishpancheta/api-gateway-service/pkg/listings"
	"github.com/pandishpancheta/api-gateway-service/pkg/order"
)

func main() {

	cfg := config.LoadConfig()

	r := mux.NewRouter()
	r.Use(config.AccessControlMiddleware)

	authSvc := auth.RegisterRouters(r, cfg)

	listings.RegisterRouters(r, cfg, authSvc.AuthClient)

	order.RegisterRouters(r, cfg, authSvc.AuthClient)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
