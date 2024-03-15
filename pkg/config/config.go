package config

import (
	"os"
)

type Config struct {
	Port                   string `mapstructure:"PORT" default:"3001"`
	ListingsServiceAddress string
	AuthServiceAddress     string
	OrderServiceAddress    string
}

func LoadConfig() *Config {
	return &Config{
		Port:                   os.Getenv("PORT"),
		ListingsServiceAddress: os.Getenv("LISTINGS_SERVICE_ADDRESS"),
		AuthServiceAddress:     os.Getenv("AUTH_SERVICE_ADDRESS"),
		OrderServiceAddress:    os.Getenv("ORDER_SERVICE_ADDRESS"),
	}
}
