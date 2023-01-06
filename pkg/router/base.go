package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetUp() *fiber.App {
	f := fiber.New()
	user := f.Group("/user")
	NewUserRouter(user).UserRouters()

	product := f.Group("/prduct")
	NewProductRouter(product).ProductRouters()
	return f
}
