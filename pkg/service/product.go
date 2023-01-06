package service

import (
	"github.com/eminoz/go-api/pkg/core/utilities"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	CreateProduct(ctx *fiber.Ctx) *utilities.ResultOfSuccessData
}
type productService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(p repository.ProductRepository) ProductService {
	return &productService{
		ProductRepository: p,
	}
}
func (p productService) CreateProduct(ctx *fiber.Ctx) *utilities.ResultOfSuccessData {
	product := new(model.Product)
	ctx.BodyParser(product)

	createdProduct := p.ProductRepository.CreateProduct(*product)
	return utilities.SuccessDataResult("product Created", createdProduct)
}
