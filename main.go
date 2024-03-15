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

	authSvc := auth.RegisterRouters(r, cfg)

	listings.RegisterRouters(r, cfg, authSvc.AuthClient)

	order.RegisterRouters(r, cfg, authSvc.AuthClient)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
