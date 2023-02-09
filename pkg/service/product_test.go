package service

import (
	"testing"

	"github.com/eminoz/go-api/pkg/internal/mocks/repomocks"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateProduct(t *testing.T) {
	mockProdoct := new(repomocks.ProductRepository)
	prod := model.Product{ProductName: "elma", Quantity: 1, Price: 23, Description: "dene bunu"}
	id := primitive.ObjectID{123}
	dto := model.ProductDto{ProductName: "elma", Quantity: 1, Price: 23, Description: "dene bunu", ID: id}
	mockProdoct.On("CreateProduct", prod).Return(dto)
	s := NewProductService(mockProdoct)
	z := s.CreateProduct(prod)
	assert.Equal(t, prod.ProductName, z.ProductName)
}
