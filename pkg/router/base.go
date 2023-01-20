package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func SetUp() *fiber.App {
	f := fiber.New()
	//monitoring
	f.Get("/monitoring", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
	user := f.Group("/user")
	NewUserRouter(user).UserRouters()

	product := f.Group("/product")
	NewProductRouter(product).ProductRouters()
	return f
}
