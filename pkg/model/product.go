package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ProductName string `json:"productName"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
}
type ProductDto struct {
	ProductName string             `json:"productName"`
	Quantity    int64              `json:"quantity"`
	Price       int64              `json:"price"`
	Description string             `json:"description"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
