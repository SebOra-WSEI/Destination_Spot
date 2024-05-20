package token

import (
	"fmt"
	"github.com/SebastianOraczek/auth/internal/env"
	"github.com/SebastianOraczek/auth/internal/model"
	"github.com/SebastianOraczek/auth/internal/response"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

var expireTime = (time.Hour * 1).Milliseconds()

func Create(user model.User) (string, error) {
	expAt := time.Now().UnixMilli() + expireTime
	secretKey := env.GetEnvVariableByName(env.JwtSecretKey)

	claims := model.Claims{
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

func Verify(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, response.ErrAuthTokenNotFound
	}

	tokenSlice := strings.Split(authHeader, " ")

	if len(tokenSlice) != 2 {
		return nil, response.ErrAuthTokenIncorrectFormat
	}

	if tokenSlice[0] != "Bearer" {
		return nil, response.ErrAuthTokenIncorrectFormat
	}

	token, err := jwt.ParseWithClaims(
		tokenSlice[1], jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetEnvVariableByName(env.JwtSecretKey)), nil
		},
	)

	if err != nil {
		fmt.Println("Parsing token error:", err.Error())
		return nil, response.ErrAuthTokenIncorrectFormat
	}

	expTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		fmt.Println("Problem with getting expiration time:", err.Error())
		return nil, response.ErrInternalServer
	}

	token.Valid = expTime.Unix()-time.Now().UnixMilli() > 0

	if !token.Valid {
		return nil, response.ErrTokenExpired
	}

	return token, nil
}
