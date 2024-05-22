package token

import (
	"fmt"
	"github.com/SebOra-WSEI/Destination_spot/shared/env"
	"github.com/SebOra-WSEI/Destination_spot/shared/message"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
	"time"
)

func Verify(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, message.ErrAuthTokenNotFound
	}

	tokenSlice := strings.Split(authHeader, " ")

	if len(tokenSlice) != 2 {
		return nil, message.ErrAuthTokenIncorrectFormat
	}

	if tokenSlice[0] != "Bearer" {
		return nil, message.ErrAuthTokenIncorrectFormat
	}

	jwtSecretKey, err := env.GetEnvVariableByName(env.JwtSecretKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	token, err := jwt.ParseWithClaims(
		tokenSlice[1], jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)

	if err != nil {
		fmt.Println("Parsing token error:", err.Error())
		return nil, message.ErrAuthTokenIncorrectFormat
	}

	expTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		fmt.Println("Problem with getting expiration time:", err.Error())
		return nil, message.ErrInternalServer
	}

	token.Valid = expTime.Unix()-time.Now().UnixMilli() > 0

	if !token.Valid {
		return nil, message.ErrTokenExpired
	}

	return token, nil
}
