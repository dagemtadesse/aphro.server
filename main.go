package main

import (
	"log"

	"aphro.web/controller/auth"
	"aphro.web/controller/files"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Use(auth.AuthMiddleWare)

	app.Post("/login", auth.Login)

	app.Post("/register", auth.Register)

	app.Post("/profile-pic/", files.PostProfilePic)

	log.Fatalln(app.Listen(":8080"))
}
