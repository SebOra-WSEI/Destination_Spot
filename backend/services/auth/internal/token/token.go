package token

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/env"
	"github.com/SebOra-WSEI/Destination_spot/shared/model"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

var expireTime = (time.Hour * 1).Milliseconds()

type Claims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}

func Create(user model.User) (string, error) {
	expAt := time.Now().UnixMilli() + expireTime
	secretKey, err := env.GetEnvVariableByName(env.JwtSecretKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	claims := Claims{
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expAt, 0)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("")
	}

	return token, nil
}
