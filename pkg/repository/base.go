package repository

import (
	"github.com/eminoz/go-api/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	DB *mongo.Database
	Cl *mongo.Collection
}
type ProductCollection struct {
	DB *mongo.Database
	Cl *mongo.Collection
}

var database = db.GetDatabase()

func UserCollectionSetting() *UserCollection {

	return &UserCollection{
		DB: database,
		Cl: database.Collection("user"),
	}
}

func ProductCollectionSetting() *ProductCollection {
	return &ProductCollection{
		DB: database,
		Cl: database.Collection("product"),
	}
}
