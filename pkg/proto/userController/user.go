package usercontroller

import (
	"context"
	"github.com/eminoz/go-api/pkg/model"
	proto "github.com/eminoz/go-api/pkg/proto/pb"
	"github.com/eminoz/go-api/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProto struct {
	UserProtos service.UserService
}

func (u UserProto) CreateUser(ctx context.Context, user *proto.User) (*proto.AuthDto, error) {
	//TODO implement me
	usr := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		ID:       primitive.ObjectID{},
	}
	r, err := u.UserProtos.CreateUser(&usr)
	if err != nil {
		return nil, nil
	}
	var a = proto.AuthDto{
		UserDto: []*proto.UserDto{{ID: string(r.ID.String()), Name: r.Name, Email: r.Email}},
		Token:   r.Token,
	}
	return &a, nil
}

func (u UserProto) GetUser(ctx context.Context, id *proto.UserID) (*proto.UserDto, error) {
	user, err := u.UserProtos.GetUser(id.UserId)
	if err != nil {
		return nil, err
	}
	usr := proto.UserDto{ID: user.ID.String(), Name: user.Name, Email: user.Email}
	return &usr, nil
}

func (u UserProto) DeleteUserById(ctx context.Context, id *proto.UserID) (*proto.ResponseMessage, error) {
	msg := u.UserProtos.DeleteUserById(id.UserId)
	return &proto.ResponseMessage{Message: msg}, nil
}

func (u UserProto) UpdateUserById(ctx context.Context, user *proto.User) (*proto.ResponseMessage, error) {
	id := user.ID
	m := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		ID:       primitive.ObjectID{},
	}
	updateUserById := u.UserProtos.UpdateUserById(id, &m)
	return &proto.ResponseMessage{Message: updateUserById}, nil
}
func (u UserProto) GetAllUser(ctx context.Context, null *proto.Null) (*proto.ResponseUsersDto, error) {

	allUser := u.UserProtos.GetAllUser()
	usersDto := proto.ResponseUsersDto{}
	for i, dto := range allUser {
		usersDto.UDto[i].Email = dto.Email
		usersDto.UDto[i].Name = dto.Name
		usersDto.UDto[i].ID = dto.ID.String()
	}
	dto := proto.ResponseUsersDto{UDto: usersDto.UDto}
	return &dto, nil

}
func (u UserProto) SignIn(ctx context.Context, authentication *proto.Authentication) (*proto.AuthDto, error) {
	auth := model.Authentication{
		Email:    authentication.Email,
		Password: authentication.Password,
	}
	r, b := u.UserProtos.SignIn(&auth)
	if b != nil {
		return nil, b
	}
	dtos := []*proto.UserDto{{ID: string(r.ID.String()), Name: r.Name, Email: r.Email}}
	var a = proto.AuthDto{
		UserDto: dtos,
		Token:   r.Token,
	}
	return &a, nil
}
