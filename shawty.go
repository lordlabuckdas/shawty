package main

import (
	"context"
	"fmt"
	// "fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	ShortURL string `bson:"short_url"`
	LongURL  string `bson:"long_url"`
}

var serverURL string = "http://localhost:3000"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	coll := client.Database("shawty").Collection("urls")
	engine := html.New("./web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/url/:shortURL", func(c *fiber.Ctx) error {
		shortURL := c.Params("shortURL")
		var res bson.M
		err := coll.FindOne(c.Context(), bson.D{primitive.E{
			Key:   "short_url",
			Value: shortURL,
		}}).Decode(&res)
		fmt.Println(res["long_url"])
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Println("no matching results")
			} else {
				log.Fatalln(err)
			}
			return nil
		}
		return c.Redirect(res["long_url"].(string))
	})
	app.Post("/", func(c *fiber.Ctx) error {
		longURL := c.FormValue("longURL")
		shortURL := uuid.New().String()[:7]
		val, err := coll.InsertOne(c.Context(), bson.D{
			primitive.E{Key: "long_url", Value: longURL},
			primitive.E{Key: "short_url", Value: shortURL},
		})
		fmt.Println(val)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Render("index", fiber.Map{
			"longURL":  longURL,
			"shortURL": serverURL + "/url/" + shortURL,
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Listen(":3000")
}
