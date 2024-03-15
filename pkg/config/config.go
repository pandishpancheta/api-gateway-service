package config

import "github.com/spf13/viper"

type Config struct {
	Port                   string `mapstructure:"PORT" default:"3001"`
	ListingsServiceAddress string
	AuthServiceAddress     string
	OrderServiceAddress    string
	UserServiceAddress     string
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("envs/")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
