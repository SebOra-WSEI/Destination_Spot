package token

import (
	"fmt"
	"github.com/SebOra-WSEI/auth/internal/env"
	"github.com/SebOra-WSEI/auth/internal/model"
	"github.com/golang-jwt/jwt/v5"
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
