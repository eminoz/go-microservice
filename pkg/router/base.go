package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetUp() *fiber.App {
	f := fiber.New()
	user := f.Group("/user")
	a := UserRouter{Route: user}
	a.UserRouters()
	return f
}
