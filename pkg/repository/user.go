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
	DeleteUserById(id string) string
	GetUserByEmailForAuth(email string) model.User
	UpdateUserById(id string, user model.User) (bool, string)
}

func (u UserCollection) DeleteUserById(id string) string {
	var ctx context.Context
	userId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userId}}
	deleteOne, _ := u.Cl.DeleteOne(ctx, filter)
	if deleteOne.DeletedCount == 1 {
		return "user delete succesfuly"
	}
	return "user did not delete"
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
func (u UserCollection) GetUserByEmailForAuth(email string) model.User {
	var ctx context.Context
	filter := bson.D{{Key: "email", Value: email}}
	var user model.User
	u.Cl.FindOne(ctx, filter).Decode(&user)
	return user
}
func (u UserCollection) UpdateUserById(id string, user model.User) (bool, string) {
	var ctx context.Context
	userID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: user.Name}}}}
	updateOne, _ := u.Cl.UpdateOne(ctx, filter, update)
	print(updateOne.ModifiedCount)
	if updateOne.ModifiedCount > 0 {
		return true, "user updated"
	}
	return false, "user did not updated"
}
