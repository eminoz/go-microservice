package jwt

import (
	"fmt"
	"time"

	"github.com/eminoz/go-api/pkg/config"
	"github.com/golang-jwt/jwt"
)

type AuthJwt interface {
	GenerateJWT(email string, role string) (string, error)
}

//go:generate mockgen -destination=../mocks/Auth/mockUserAuth.go -package=jwt  github.com/eminoz/go-advanced-microservice/security/jwt IToken

type authJwt struct{}

func NewAuthJwt() AuthJwt {
	return &authJwt{}
}
func (a authJwt) GenerateJWT(email string, role string) (string, error) {
	secretKey := config.GetConfig().AppSecret
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	fmt.Print(tokenString)
	return tokenString, nil
}
