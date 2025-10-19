package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port        string
	Environment string
}

/**
 * Get environment variable or return default value
 */
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (e *Env) Validate() error {
	if e.Port == "" {
		return fmt.Errorf("PORT is required")
	}

	if e.Environment == "" {
		return fmt.Errorf("ENVIRONMENT is required")
	}

	if e.Environment != "development" && e.Environment != "production" && e.Environment != "staging" {
		return fmt.Errorf("ENVIRONMENT must be one of 'development', 'production', or 'staging'")
	}

	return nil
}

/**
 * Load environment variables from .env file
 */
func LoadEnv() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	env := &Env{
		Port:        getEnv("PORT", "8000"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	if err := env.Validate(); err != nil {
		log.Fatal("Invalid env config", err)
	}

	return env
}

func (e *Env) IsDevelopment() bool {
	return e.Environment == "development"
}

func (e *Env) IsProduction() bool {
	return e.Environment == "production"
}

func (e *Env) IsStaging() bool {
	return e.Environment == "staging"
}
