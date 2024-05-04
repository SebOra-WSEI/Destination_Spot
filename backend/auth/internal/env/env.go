package env

import (
	"log"
	"os"
)

const ConnectionStringEnvName string = "CONNECTION_STRING"

func GetEnvVariableByName(name string) string {
	env := os.Getenv(name)

	if env == "" {
		log.Fatalf("Environment variable %v is empty", name)
	}

	return env
}
