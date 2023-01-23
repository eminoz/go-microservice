package main

import (
	"context"
	"fmt"
	proto "github.com/eminoz/grpc-deneme/proto/pb"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"log"
)

type user struct {
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewUserServiceClient(conn)
	// g := gin.Default()
	app := fiber.New()
	app.Use(logger.New())
	app.Post("/createuser", func(ctx *fiber.Ctx) error {

		user := new(proto.User)
		ctx.BodyParser(&user)
		fmt.Println(user)
		a, b := client.CreateUser(context.Background(), user)
		fmt.Println(a.UserDto)
		if b != nil {
			return ctx.JSON(b)
		}
		return ctx.JSON(a)
	})
	app.Get("/getuser/:id", func(ctx *fiber.Ctx) error {
		userId := ctx.Params("id")
		id := proto.UserID{UserId: userId}
		a, b := client.GetUser(context.Background(), &id)
		if b != nil {
			return ctx.JSON(b)
		}
		return ctx.JSON(a)
	})
	log.Fatal(app.Listen(":3000"))
}
