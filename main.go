package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/s809616134/hotel-backend/api"
	"github.com/s809616134/hotel-backend/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const userColl = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection((userColl))

	user := types.User{
		FirstName: "james",
		LastName:  "dude",
	}
	res, err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

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
