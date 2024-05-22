package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	ConnectionString string = "CONNECTION_STRING"
	Domain                  = "DOMAIN"
	JwtSecretKey            = "JWT_SECRET_KEY"
)

func Load(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}

func GetEnvVariableByName(name string) (string, error) {
	env := os.Getenv(name)
	if env == "" {
		return "", fmt.Errorf("environment variable %v is empty", name)
	}

	return env, nil
}
