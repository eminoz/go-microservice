package repository

import (
	"context"

	"github.com/eminoz/go-api/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(user *model.User) model.UserDto
	GetUserByID(id string) model.UserDto
}

func (u UserCollection) CreateUser(user *model.User) model.UserDto {
	var ctx context.Context
	response, err := u.Cl.InsertOne(ctx, user)
	if err != nil {
		return model.UserDto{}
	}
	id := response.InsertedID

	filter := bson.D{{Key: "_id", Value: id}}
	var userDto model.UserDto
	u.Cl.FindOne(ctx, filter).Decode(&userDto)
	return userDto

}
func (u UserCollection) GetUserByID(id string) model.UserDto {
	var ctx context.Context
	userID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userID}}
	var userDto model.UserDto
	u.Cl.FindOne(ctx, filter).Decode(&userDto)
	return userDto
}
