package service

import (
	"github.com/eminoz/go-api/pkg/broker"
	"github.com/eminoz/go-api/pkg/cache"
	"github.com/eminoz/go-api/pkg/core/utilities"

	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/security/encryption"
	"github.com/eminoz/go-api/pkg/security/jwt"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError)
	GetUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError)
	DeleteUserById(ctx *fiber.Ctx) *utilities.ResultSuccess
	UpdateUserById(ctx *fiber.Ctx) *utilities.ResultSuccess
	GetAllUser() *utilities.DataResult
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

func (u userService) CreateUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError) {
	user := new(model.User)
	ctx.BodyParser(user)
	// userInDB := u.UserRepository.GetUserByEmailForAuth(user.Email)
	// if userInDB.Email != "" {
	// 	return nil, utilities.ErrorResult("user already created ")
	// }
	u.UserBroker.CreatedUser(*user) //send user to createUser queue
	bycripted, err := u.Encryption.GenerateHashPassword(user.Password)
	if err != nil {
		return nil, utilities.ErrorResult("")
	}
	user.Password = bycripted //give user model the bycripted password
	responseUser := u.UserRepository.CreateUser(user)
	token, _ := u.Authentication.CreateToken(user.Email, user.Password)

	u.UserCache.SaveUserEmailByID(responseUser.ID.Hex(), responseUser.Email) //save user email by id in redis

	userDto := model.AuthDto{UserDto: responseUser, Token: token.TokenString}
	return utilities.SuccessDataResult("user created", userDto), nil

}
func (u userService) GetUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError) {
	userId := ctx.Params("id")
	user := u.UserRepository.GetUserByID(userId)
	return utilities.SuccessDataResult("user", user), nil
}
func (u userService) DeleteUserById(ctx *fiber.Ctx) *utilities.ResultSuccess {
	userID := ctx.Params("id")
	deletedUser := u.UserRepository.DeleteUserById(userID)
	return utilities.SuccessResult(deletedUser)
}
func (u userService) UpdateUserById(ctx *fiber.Ctx) *utilities.ResultSuccess {
	userID := ctx.Params("id")
	user := new(model.User)

	ctx.BodyParser(&user)

	_, msg := u.UserRepository.UpdateUserById(userID, *user)
	return utilities.SuccessResult(msg)
}
func (u userService) GetAllUser() *utilities.DataResult {
	allUsers := u.UserRepository.GetAllUser()
	return utilities.SuccessDataResult("all user", allUsers)
}
