package service

import (
	"github.com/eminoz/go-api/pkg/core/utilities"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	CreateUser(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError)
}
type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(u repository.UserRepository) UserService {
	return &userService{
		UserRepository: u,
	}
}
func (u userService) CreateUser(ctx *fiber.Ctx) (*utilities.ResultOfSuccessData, *utilities.ResultError) {

	user := new(model.User)
	ctx.BodyParser(user)
	responseUser := u.UserRepository.CreateUser(user)
	return utilities.SuccessDataResult("user created", responseUser), nil

}
