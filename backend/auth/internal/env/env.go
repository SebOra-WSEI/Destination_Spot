package env

import (
	"log"
	"os"
)

const (
	ConnectionString string = "CONNECTION_STRING"
	Domain                  = "DOMAIN"
)

func GetEnvVariableByName(name string) string {
	env := os.Getenv(name)
	if env == "" {
		log.Fatalf("Environment variable %v is empty", name)
	}

	return env
}
