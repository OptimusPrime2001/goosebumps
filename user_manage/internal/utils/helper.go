package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if value == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return value
}
func LoadEnv() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current working directory: %v", err)
	}
	envPath := filepath.Join(cwd, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Println("failed to load .env:", err)
	} else {
		log.Printf("✅ Loaded .env file from %s successfully", envPath)
	}
}
