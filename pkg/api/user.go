package api

import (
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type UserApi interface {
	CreateUser(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	DeleteUserById(ctx *fiber.Ctx) error
	UpdateUserById(ctx *fiber.Ctx) error
	GetAllUser(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
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
	user := new(model.User)
	ctx.BodyParser(user)
	response, err := u.UserService.CreateUser(user)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(response)
}
func (u userApi) GetUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := u.UserService.GetUser(userId)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(user)

}
func (u userApi) DeleteUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	deletedUser := u.UserService.DeleteUserById(userID)
	return ctx.JSON(deletedUser)

}
func (u userApi) UpdateUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	user := new(model.User)

	ctx.BodyParser(&user)
	response := u.UserService.UpdateUserById(userID, user)
	return ctx.JSON(response)
}
func (u userApi) GetAllUser(ctx *fiber.Ctx) error {

	allUsers := u.UserService.GetAllUser()

	return ctx.JSON(allUsers)
}
func (u userApi) SignIn(ctx *fiber.Ctx) error {
	auth := new(model.Authentication)
	ctx.BodyParser(auth)
	user, err := u.UserService.SignIn(auth)
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.JSON(user)
}
