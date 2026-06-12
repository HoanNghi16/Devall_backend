package auth

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID uint
	Role   string
 	jwt.RegisteredClaims
}