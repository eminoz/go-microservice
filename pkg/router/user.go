package router

import "github.com/gofiber/fiber/v2"

type UserRouter struct {
	Route fiber.Router
}

func (u UserRouter) UserRouters() {
	di := base{}
	var user = di.UserDI()
	u.Route.Post("/create", user.CreateUser)
	u.Route.Get("/getUser/:id", user.GetUser)
}
