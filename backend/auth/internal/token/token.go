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

func CreateToken(user model.User) (string, error) {
	expAt := time.Now().UnixMilli() + expireTime
	secretKey := env.GetEnvVariableByName(env.JwtSecretKey)

	claims := model.Claims{
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expAt, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("")
	}

	return tokenString, nil
}

func Verify(authHeader string) (*jwt.Token, error) {
	if authHeader == "" {
		return nil, fmt.Errorf(response.AuthTokenNotFoundErrMsg)
	}

	tokenSlice := strings.Split(authHeader, " ")

	if len(tokenSlice) != 2 {
		return nil, fmt.Errorf(response.AuthTokenIncorrectFormatErrMsg)
	}

	if tokenSlice[0] != "Bearer" {
		return nil, fmt.Errorf(response.AuthTokenIncorrectFormatErrMsg)
	}

	token, err := jwt.ParseWithClaims(
		tokenSlice[1], jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetEnvVariableByName(env.JwtSecretKey)), nil
		},
	)

	if err != nil {
		fmt.Println("Parsing token error:", err.Error())
		return nil, fmt.Errorf(response.AuthTokenIncorrectFormatErrMsg)
	}

	expTime, err := token.Claims.GetExpirationTime()
	if err != nil {
		fmt.Println("Problem with getting expiration time:", err.Error())
		return nil, fmt.Errorf(response.InternalServerErrMsg)
	}

	token.Valid = expTime.Unix()-time.Now().UnixMilli() > 0

	if !token.Valid {
		return nil, fmt.Errorf(response.TokenExpiredErrMsg)
	}

	return token, nil
}
