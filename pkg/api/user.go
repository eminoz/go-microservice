package api

import (
	"github.com/eminoz/go-api/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(ctx *fiber.Ctx) error
}
type userApi struct {
	UserService service.UserService
}

func NewUserApi(s service.UserService) UserApi {
	return &userApi{
		UserService: s,
	}
}

func (u userApi) CreateUser(ctx *fiber.Ctx) error {
	response, err := u.UserService.CreateUser(ctx)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(response)
}
