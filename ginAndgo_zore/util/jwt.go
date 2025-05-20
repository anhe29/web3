package util

import (
	"ginAndgo_zore/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Id       int
	Username string
	Email    string
	jwt.StandardClaims
}

func GenerateToken(user model.User, secret string) (string, error) {
	claims := &Claim{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString, secret string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
