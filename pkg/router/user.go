package router

import "github.com/gofiber/fiber/v2"

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
	u.Route.Post("/create", user.CreateUser)
	u.Route.Get("/getUser/:id", user.GetUser)
	u.Route.Get("/deleteUser/:id", user.DeleteUserById)
}
