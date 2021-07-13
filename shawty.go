package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URL struct {
	ShortURL string `json:"short_url" bson:"short_url"`
	LongURL  string `json:"long_url" bson:"long_url"`
}

var (
	serverURL string = "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT")
	mongoDBURI string = os.Getenv("MONGO_URI")
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(mongoDBURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	coll := client.Database("shawty").Collection("urls")
	engine := html.New("./web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Static("/", "./web")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})
	app.Post("/", func(c *fiber.Ctx) error {
		longURL, err := url.Parse(c.FormValue("longURL"))
		// log.Println("URL received: "+ longURL.String())
		if err != nil {
			return c.Status(400).Render("home", fiber.Map{
				"error": "400 - Invalid URL",
			})
		}
		if len(longURL.Scheme) == 0 {
			longURL.Scheme = "http"
		}
		// log.Println("Final URL: "+ longURL.String())
		url := URL{
			ShortURL: uuid.New().String()[:6],
			LongURL: longURL.String(),
		}
		bsonURL, err := bson.Marshal(url)
		if err != nil {
			return c.Status(500).Render("home", fiber.Map{
				"error": "500 - Error marshalling to BSON",
			})
		}
		// log.Println("Marshalled")
		_, err = coll.InsertOne(c.Context(), bsonURL)
		if err != nil {
			return c.Status(500).Render("home", fiber.Map{
				"error": "500 - Error adding URL to the DB",
			})
		}
		// log.Println("Added to DB")
		return c.Render("home", fiber.Map{
			"longURL":  longURL,
			"shortURL": serverURL + "/" + url.ShortURL,
		})
	})
	app.Get("/:shortURL", func(c *fiber.Ctx) error {
		shortURL := c.Params("shortURL")
		// log.Println("Route URL: " + shortURL)
		var url URL
		err := coll.FindOne(c.Context(), bson.D{primitive.E{
			Key:   "short_url",
			Value: shortURL,
		}}).Decode(&url)
		// log.Println("Executed Search")
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(404).Render("home", fiber.Map{
					"error": "404 - Page Not Found",
				})
			} else {
				return c.Status(400).Render("home", fiber.Map{
					"error": "400 - Bad Request",
				})
			}
		}
		// log.Println("Successful search")
		return c.Status(302).Redirect(url.LongURL)
	})
	log.Fatalln(app.Listen(serverURL[7:]))
}
