package router

import "github.com/gofiber/fiber/v2"

type ProductRouter struct {
	Router fiber.Router
}

func NewProductRouter(r fiber.Router) *ProductRouter {
	return &ProductRouter{
		Router: r,
	}
}

func (p ProductRouter) ProductRouters() {
	r := p.Router
	r.Post("/create")
}
