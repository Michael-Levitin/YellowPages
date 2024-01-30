package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}

func Init() {
	// загружаем данные из .env файла в систему
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	log.Println("loaded env values")
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		DbHost:     getEnv("DB_HOST", ""),
		DbPort:     getEnv("DB_PORT", ""),
		DbName:     getEnv("DB_NAME", ""),
		DbUsername: getEnv("DB_USERNAME", ""),
		DbPassword: getEnv("DB_PASSWORD", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	var value string
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if value == "" && defaultVal == "" {
		log.Fatal(key, " value not found")
	}
	return defaultVal
}
