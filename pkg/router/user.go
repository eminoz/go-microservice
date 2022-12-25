package router

import "github.com/gofiber/fiber/v2"

type UserRouter struct {
	Route fiber.Router
}

func (u UserRouter) UserRouters() {
	u.Route.Get("/get", func(c *fiber.Ctx) error {
		return c.JSON("hello there")
	})

}
