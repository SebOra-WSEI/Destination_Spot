package model

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Email string
	Role  string
	jwt.RegisteredClaims
}
