package service

import (
	"github.com/eminoz/go-api/pkg/broker"
	"github.com/eminoz/go-api/pkg/cache"
	"github.com/eminoz/go-api/pkg/core/utilities"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/security/encryption"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError)
	GetUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError)
}
type userService struct {
	UserRepository repository.UserRepository
	UserBroker     broker.User
	UserCache      cache.UserCache
	Encryption     encryption.Encryption
}

func NewUserService(u repository.UserRepository, b broker.User, c cache.UserCache, e encryption.Encryption) UserService {
	return &userService{
		UserRepository: u,
		UserBroker:     b,
		UserCache:      c,
		Encryption:     e,
	}
}
func (u userService) CreateUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError) {
	user := new(model.User)
	ctx.BodyParser(user)
	u.UserBroker.CreatedUser(*user) //send user to createUser queue
	bycripted, err := u.Encryption.GenerateHashPassword(user.Password)
	if err != nil {
		return nil, utilities.ErrorResult("")
	}
	user.Password = bycripted //give user model the bycripted password
	responseUser := u.UserRepository.CreateUser(user)
	u.UserCache.SaveUserEmailByID(responseUser.ID.Hex(), responseUser.Email) //save user email by id in redis
	return utilities.SuccessDataResult("user created", responseUser), nil

}
func (u userService) GetUser(ctx *fiber.Ctx) (*utilities.DataResult, *utilities.ResultError) {
	userId := ctx.Params("id")
	user := u.UserRepository.GetUserByID(userId)
	return utilities.SuccessDataResult("user", user), nil
}
