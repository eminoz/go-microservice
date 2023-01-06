package repository

import (
	"context"

	"github.com/eminoz/go-api/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository interface {
	CreateProduct(product model.Product) model.ProductDto
	GetProductByID(id string)
}

func (p ProductCollection) CreateProduct(product model.Product) model.ProductDto {
	var ctx context.Context

	InsertOne, _ := p.Cl.InsertOne(ctx, product)
	filter := bson.D{{Key: "_id", Value: InsertOne.InsertedID}}
	var productDto model.ProductDto
	p.Cl.FindOne(ctx, filter).Decode(&productDto)
	return productDto

}

func (p ProductCollection) GetProductByID(id string) {
	panic("not implemented") // TODO: Implement
}
