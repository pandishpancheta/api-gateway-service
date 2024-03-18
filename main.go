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
	"github.com/rs/cors"
)

func main() {

	cfg := config.LoadConfig()

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://*.vercel.app", "http://localhost:3000/", "https://*.vercel.app/", "http://*.kaloyan.tech/", "https://*.kaloyan.tech", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	authSvc := auth.RegisterRouters(r, cfg)

	listings.RegisterRouters(r, cfg, authSvc.AuthClient)

	order.RegisterRouters(r, cfg, authSvc.AuthClient)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	handler := c.Handler(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))
}
