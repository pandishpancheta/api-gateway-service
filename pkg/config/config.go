package config

import (
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Port                   string `mapstructure:"PORT" default:"3001"`
	ListingsServiceAddress string
	AuthServiceAddress     string
	OrderServiceAddress    string
}

func LoadConfig() *Config {
	return &Config{
		Port:                   strings.TrimSuffix(os.Getenv("PORT"), "\n"),
		ListingsServiceAddress: strings.TrimSuffix(os.Getenv("LISTINGS_SERVICE_ADDRESS"), "\n"),
		AuthServiceAddress:     strings.TrimSuffix(os.Getenv("AUTH_SERVICE_ADDRESS"), "\n"),
		OrderServiceAddress:    strings.TrimSuffix(os.Getenv("ORDERS_SERVICE_ADDRESS"), "\n"),
	}
}

func AccessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
