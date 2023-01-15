package jwt

import (
	"time"

	"github.com/eminoz/go-api/pkg/config"
	"github.com/eminoz/go-api/pkg/core/utilities"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/golang-jwt/jwt"
)

type AuthJwt interface {
	GenerateJWT(email string, role string) (string, error)
	CreateToken(email string, password string) (model.Token, *utilities.ResultError)
}

//go:generate mockgen -destination=../mocks/Auth/mockUserAuth.go -package=jwt  github.com/eminoz/go-advanced-microservice/security/jwt IToken

type authJwt struct {
	UserRepository repository.UserRepository
}

func NewAuthJwt(UserRepository repository.UserRepository) AuthJwt {
	return &authJwt{
		UserRepository: UserRepository,
	}
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
	return tokenString, nil
}
func (a authJwt) CreateToken(email string, password string) (model.Token, *utilities.ResultError) {
	user := a.UserRepository.GetUserByEmailForAuth(email)

	// checkPasswordHash := u.Encryption.CheckPasswordHash(password, user.Password)
	// if !checkPasswordHash {
	// 	return model.Token{}, utilities.ErrorResult("password is incorrect")
	// }
	generateJWT, err := a.GenerateJWT(user.Email, user.Role)

	if err != nil {
		return model.Token{}, utilities.ErrorResult("did not generate token")
	}
	var token model.Token
	token.Email = user.Email
	token.Role = user.Role
	token.ID = user.ID
	token.TokenString = generateJWT
	return token, nil
}
