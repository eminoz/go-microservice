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

func (u UserProto) DeleteUserById(ctx context.Context, id *proto.UserID) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserProto) UpdateUserById(ctx context.Context, user *proto.User) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserProto) GetAllUser(ctx context.Context, null *proto.Null) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserProto) SignIn(ctx context.Context, authentication *proto.Authentication) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}
