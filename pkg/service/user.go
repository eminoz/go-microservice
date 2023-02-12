package service

import (
	"errors"

	"github.com/eminoz/go-api/pkg/broker"
	"github.com/eminoz/go-api/pkg/cache"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/security/encryption"
	"github.com/eminoz/go-api/pkg/security/jwt"
)

type UserService interface {
	CreateUser(user *model.User) (model.AuthDto, error)
	GetUser(userId string) (model.UserDto, error)
	DeleteUserById(userID string) string
	UpdateUserById(userID string, user *model.User) string
	GetAllUser() []model.UserDto
	SignIn(auth *model.Authentication) (model.AuthDto, error)
}
type userService struct {
	UserRepository repository.UserRepository
	UserBroker     broker.User
	UserCache      cache.UserCache
	Encryption     encryption.Encryption
	Authentication jwt.AuthJwt
}

func NewUserService(u repository.UserRepository, b broker.User, c cache.UserCache, e encryption.Encryption, a jwt.AuthJwt) UserService {
	return &userService{
		UserRepository: u,
		UserBroker:     b,
		UserCache:      c,
		Encryption:     e,
		Authentication: a,
	}
}

func (u userService) CreateUser(user *model.User) (model.AuthDto, error) {

	// userInDB := u.UserRepository.GetUserByEmailForAuth(user.Email)
	// if userInDB.Email != "" {
	// 	return model.AuthDto{}, nil
	// }
	u.UserBroker.CreatedUser(*user) //send user to createUser queue
	bycripted, err := u.Encryption.GenerateHashPassword(user.Password)
	if err != nil {
		err := errors.New("could not generate hash code")
		return model.AuthDto{}, err
	}
	user.Password = bycripted //give user model the bycripted password
	responseUser := u.UserRepository.CreateUser(user)
	token, _ := u.Authentication.CreateToken(user.Email, user.Password)

	userDto := model.AuthDto{UserDto: responseUser, Token: token.TokenString}
	u.UserCache.SaveUserByID(responseUser.ID.Hex(), responseUser) //save user email by id in redis
	return userDto, nil
}
func (u userService) SignIn(auth *model.Authentication) (model.AuthDto, error) {

	token, err := u.Authentication.CreateToken(auth.Email, auth.Password)
	if err != nil {
		return model.AuthDto{}, err
	}
	responseUser := u.UserRepository.GetUserByEmail(auth.Email)
	userDto := model.AuthDto{UserDto: responseUser, Token: token.TokenString}
	return userDto, nil
}
func (u userService) GetUser(userId string) (model.UserDto, error) {

	redisUser := u.UserCache.GetUSerById(userId) //get user from redis
	if redisUser.Email != "" {
		return model.UserDto{ID: redisUser.ID, Name: redisUser.Name, Email: redisUser.Email}, nil
	}
	user := u.UserRepository.GetUserByID(userId)
	if user.Email == "" {
		return model.UserDto{}, nil
	}
	u.UserCache.SaveUserByID(user.ID.Hex(), user) //save user by id in redis
	return user, nil
}
func (u userService) DeleteUserById(userID string) string {

	deletedUser := u.UserRepository.DeleteUserById(userID)
	u.UserCache.DeleteUserById(userID)
	return deletedUser
}
func (u userService) UpdateUserById(userID string, user *model.User) string {

	_, msg := u.UserRepository.UpdateUserById(userID, *user)
	return msg
}
func (u userService) GetAllUser() []model.UserDto {
	allUsers := u.UserRepository.GetAllUser()
	return allUsers
}
