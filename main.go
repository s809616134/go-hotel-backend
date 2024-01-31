package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/s809616134/hotel-backend/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen addr of the api server")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("/api")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "John"})
}
