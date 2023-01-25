package main

import (
	"context"
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
		a, b := client.CreateUser(context.Background(), user)
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
	app.Get("/getalluser", func(ctx *fiber.Ctx) error {
		allUser, err2 := client.GetAllUser(context.Background(), nil)
		if err2 != nil {
			return ctx.JSON(err2)
		}
		return ctx.JSON(allUser)
	})
	app.Post("/updateuser/:id", func(ctx *fiber.Ctx) error {
		m := new(proto.User)
		ctx.BodyParser(&m)
		id := ctx.Params("id")

		m.ID = id
		byId, err2 := client.UpdateUserById(context.Background(), m)
		if err2 != nil {
			return ctx.JSON(err2)
		}
		return ctx.JSON(byId)

	})
	app.Delete("/deleteuserbyid/:id", func(ctx *fiber.Ctx) error {
		userID := ctx.Params("id")
		id := proto.UserID{UserId: userID}

		byId, err2 := client.DeleteUserById(context.Background(), &id)
		if err2 != nil {
			return ctx.JSON(err2)
		}
		return ctx.JSON(byId)
	})
	app.Post("/signin", func(ctx *fiber.Ctx) error {
		authentication := proto.Authentication{}

		ctx.BodyParser(&authentication)
		p := proto.Authentication{Email: authentication.Email, Password: authentication.Password}

		in, err2 := client.SignIn(context.Background(), &p)
		if err2 != nil {
			return ctx.JSON(err2)
		}
		return ctx.JSON(in)
	})
	log.Fatal(app.Listen(":3000"))
}
