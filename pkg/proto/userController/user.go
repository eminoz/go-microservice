package usercontroller

import (
	"context"
	proto "github.com/eminoz/go-api/pkg/proto/pb"
)

type UserProto struct{}

func (u UserProto) CreateUser(ctx context.Context, user *proto.User) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserProto) GetUser(ctx context.Context, id *proto.UserID) (*proto.Response, error) {
	//TODO implement me
	panic("implement me")
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
