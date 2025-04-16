package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server		ServerConfig
	Database	DatabaseConfig
	// JWT 		JWTConfig 
}

type ServerConfig struct {
	Port string
	Host string
}

type DatabaseConfig struct {
	Host		string
	Port		string
	User 		string
	Password	string
	Name		string
}

type JWTConfig struct {
	Secret	string
	Expiry	string
}
func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env") 
	viper.SetConfigType("env")  
	viper.AddConfigPath(".")    
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to load Config: %v", err)
	}

	config := &Config{
		Server: ServerConfig{
			Port: viper.GetString("PORT"),
			Host: viper.GetString("HOST"),
		},
		Database: DatabaseConfig{
			Host: 		viper.GetString("DB_HOST"),
			Port: 		viper.GetString("DB_PORT"),
			User: 		viper.GetString("DB_USER"),
			Password: 	viper.GetString("DB_PASSWORD"),
			Name: 		viper.GetString("DB_NAME"),
		},
	}
	return config, nil
}