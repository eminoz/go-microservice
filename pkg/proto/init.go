package proto

import (
	"fmt"
	usercontroller "github.com/eminoz/go-api/pkg/proto/userController"
	"net"

	proto "github.com/eminoz/go-api/pkg/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func BaseRPC() {
	listen, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	newServer := grpc.NewServer()
	reflection.Register(newServer)

	user := usercontroller.UserProto{}
	proto.RegisterUserServiceServer(newServer, &user)

	err = newServer.Serve(listen)
	if err != nil {
		panic(err)
	}
	fmt.Print("server stared on port 4040 ")

}
