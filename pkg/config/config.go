package config

import (
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
		Port:                   strings.TrimSpace(os.Getenv("PORT")),
		ListingsServiceAddress: strings.TrimSpace(os.Getenv("LISTINGS_SERVICE_ADDRESS")),
		AuthServiceAddress:     strings.TrimSpace(os.Getenv("AUTH_SERVICE_ADDRESS")),
		OrderServiceAddress:    strings.TrimSpace(os.Getenv("ORDER_SERVICE_ADDRESS")),
	}
}
