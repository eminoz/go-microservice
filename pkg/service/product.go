package service

import (
	"github.com/eminoz/go-api/pkg/model"
	"github.com/eminoz/go-api/pkg/repository"
)

type ProductService interface {
	CreateProduct(product model.Product) model.ProductDto
}
type productService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(p repository.ProductRepository) ProductService {
	return &productService{
		ProductRepository: p,
	}
}
func (p productService) CreateProduct(product model.Product) model.ProductDto {

	createdProduct := p.ProductRepository.CreateProduct(product)

	return createdProduct
}
