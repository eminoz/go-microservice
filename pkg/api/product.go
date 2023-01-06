package api

import (
	"github.com/eminoz/go-api/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ProductApi interface {
	CreatedProduct(ctx *fiber.Ctx) error
}
type productApi struct {
	ProductService service.ProductService
}

func NewProductApi(p service.ProductService) ProductApi {
	return &productApi{
		ProductService: p,
	}
}
func (p productApi) CreatedProduct(ctx *fiber.Ctx) error {
	c := p.ProductService.CreateProduct(ctx)
	return ctx.JSON(c)
}
