package api

import (
	"github.com/eminoz/go-api/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	DeleteUserById(ctx *fiber.Ctx) error
	UpdateUserById(ctx *fiber.Ctx) error
	GetAllUser(ctx *fiber.Ctx) error
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
func (u userApi) GetUser(ctx *fiber.Ctx) error {
	user, _ := u.UserService.GetUser(ctx)
	return ctx.JSON(user)
}
func (u userApi) DeleteUserById(ctx *fiber.Ctx) error {
	deletedUser := u.UserService.DeleteUserById(ctx)
	return ctx.JSON(deletedUser)

}
func (u userApi) UpdateUserById(ctx *fiber.Ctx) error {
	response := u.UserService.UpdateUserById(ctx)
	return ctx.JSON(response)
}
func (u userApi) GetAllUser(ctx *fiber.Ctx) error {

	allUsers := u.UserService.GetAllUser()

	return ctx.JSON(allUsers)
}
