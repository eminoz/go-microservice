package proto

import (
	"fmt"
	"github.com/eminoz/go-api/pkg/broker"
	"github.com/eminoz/go-api/pkg/cache"
	usercontroller "github.com/eminoz/go-api/pkg/proto/userController"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/security/encryption"
	"github.com/eminoz/go-api/pkg/security/jwt"
	"github.com/eminoz/go-api/pkg/service"
	"net"

	proto "github.com/eminoz/go-api/pkg/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var redis = cache.InitRedis() //connect redis
var userbroker = broker.NewUserProducer()

func BaseRPC() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	reflection.Register(newServer)
	encryption := encryption.NewUserEncription()
	userCache := cache.NewUserCache(redis) //create user redis client
	r := repository.UserCollectionSetting()
	authJwt := jwt.NewAuthJwt(r)
	s := service.NewUserService(r, userbroker, userCache, encryption, authJwt)
	user := usercontroller.UserProto{UserProtos: s}
	proto.RegisterUserServiceServer(newServer, &user)

	err = newServer.Serve(listen)
	if err != nil {
		panic(err)
	}
	fmt.Print("server stared on port 4040 ")

}
