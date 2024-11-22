package constant

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID     uint
	Device string
	jwt.RegisteredClaims
}
