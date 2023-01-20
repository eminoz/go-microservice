package router

import (
	"github.com/eminoz/go-api/pkg/middleware/security"
	"github.com/eminoz/go-api/pkg/middleware/validation"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	Route fiber.Router
}

func NewUserRouter(r fiber.Router) *UserRouter {
	return &UserRouter{
		Route: r,
	}
}
func (u UserRouter) UserRouters() {
	di := base{}
	var user = di.UserDI() //in tis function user dependencies are injected
	u.Route.Post("/create", validation.UserValidation(), user.CreateUser)
	u.Route.Post("/signin", user.SignIn)
	u.Route.Get("/getUser/:id", security.UserIsAuth(), user.GetUser)
	u.Route.Delete("/deleteUser/:id", security.UserIsAuth(), user.DeleteUserById)
	u.Route.Post("/updateUser/:id", security.UserIsAuth(), user.UpdateUserById)
	u.Route.Get("/getAllUser", security.UserIsAuth(), user.GetAllUser)
}
