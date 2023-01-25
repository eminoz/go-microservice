package api

import (
	"github.com/eminoz/go-api/pkg/core/utilities"
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
	ctx.BodyParser(&user)
	response, err := u.UserService.CreateUser(user)
	if err != nil {
		errorDataResult := utilities.ErrorDataResult(err.Error(), err)
		return ctx.JSON(errorDataResult)
	}
	successDataResult := utilities.SuccessDataResult("user created", response)
	return ctx.JSON(successDataResult)
}
func (u userApi) GetUser(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := u.UserService.GetUser(userId)
	if err != nil {
		errorDataResult := utilities.ErrorDataResult(err.Error(), err)
		return ctx.JSON(errorDataResult)
	}
	successDataResult := utilities.SuccessDataResult("user fetched successfully", user)
	return ctx.JSON(successDataResult)

}
func (u userApi) DeleteUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	deletedUser := u.UserService.DeleteUserById(userID)

	successDataResult := utilities.SuccessDataResult("user deleted successfully", deletedUser)
	return ctx.JSON(successDataResult)

}
func (u userApi) UpdateUserById(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	user := new(model.User)

	ctx.BodyParser(&user)
	response := u.UserService.UpdateUserById(userID, user)
	successDataResult := utilities.SuccessDataResult("user updated successfully", response)
	return ctx.JSON(successDataResult)
}
func (u userApi) GetAllUser(ctx *fiber.Ctx) error {

	allUsers := u.UserService.GetAllUser()

	successDataResult := utilities.SuccessDataResult("got all users successfully", allUsers)
	return ctx.JSON(successDataResult)
}
func (u userApi) SignIn(ctx *fiber.Ctx) error {
	auth := new(model.Authentication)
	ctx.BodyParser(auth)
	user, err := u.UserService.SignIn(auth)
	if err != nil {
		errorDataResult := utilities.ErrorDataResult(err.Error(), err)
		return ctx.JSON(errorDataResult)
	}
	successDataResult := utilities.SuccessDataResult("user signed in successfully", user)
	return ctx.JSON(successDataResult)
}
