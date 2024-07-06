package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/masterghost2002/videotube/internals/database"
	handlers "github.com/masterghost2002/videotube/internals/handlers/auth"
	"github.com/masterghost2002/videotube/internals/middlewares"
)

func main() {
	err := database.InitDB()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(middlewares.SQLErrorHandler)
	app.Use(middlewares.ValidatorErrorFormator)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello wordl!")
	})

	// auth router group
	authRouter := app.Group("/auth")
	authRouter.Post("/signup", handlers.SignUp)
	authRouter.Post("/signin", handlers.SignIn)
	app.Listen(":5000")
}
