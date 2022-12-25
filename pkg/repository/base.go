package repository

import (
	"github.com/eminoz/go-api/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCollection struct {
	DB *mongo.Database
	Cl *mongo.Collection
}

func UserCollectionSetting() *UserCollection {
	db := db.GetDatabase()
	return &UserCollection{
		DB: db,
		Cl: db.Collection("user"),
	}
}
