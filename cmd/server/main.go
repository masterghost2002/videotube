package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	handlers "github.com/masterghost2002/videotube/internals/handlers/auth"
	"github.com/masterghost2002/videotube/internals/registry"
)

func main() {
	err := registry.StorageInit()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello wordl!")
	})

	// auth router group
	authRouter := app.Group("/auth")
	authRouter.Post("/signup", handlers.SignUp)

	app.Listen(":5000")
}
