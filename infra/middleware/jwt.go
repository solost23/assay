package middleware

import (
	"assay/infra/constant"
	"assay/infra/global"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JWT.Key),
	}
}

// CreateToken 创建一个token
// @param claims constant.CustomClaims
// @return token string
// @return err error
func (j *JWT) CreateToken(claims constant.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
// @param tokenString string
// @return claims *constant.CustomClaims
// @return err error
func (j *JWT) ParseToken(tokenString string) (*constant.CustomClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&constant.CustomClaims{},
		keyFunc,
	)
	if err != nil {
		switch err {
		case jwt.ErrTokenMalformed:
			return nil, ErrTokenMalformed
		case jwt.ErrTokenExpired:
			// Token is expired
			return nil, ErrTokenExpired
		case jwt.ErrTokenNotValidYet:
			return nil, ErrTokenNotValidYet
		default:
			return nil, ErrTokenInvalid
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*constant.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	}
	return nil, ErrTokenInvalid
}
