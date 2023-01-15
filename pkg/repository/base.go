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

func UserCollectionSetting() *UserCollection {
	var database = db.GetDatabase()
	return &UserCollection{
		DB: database,
		Cl: database.Collection("user"),
	}
}

func ProductCollectionSetting() *ProductCollection {
	var database = db.GetDatabase()
	return &ProductCollection{
		DB: database,
		Cl: database.Collection("product"),
	}
}
