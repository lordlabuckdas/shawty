package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "github.com/go-redis/redis"

	"github.com/lordlabuckdas/shawty/handlers"
)

var (
	serverURL  string = "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT")
	mongoDBURI string = os.Getenv("MONGO_URI")
	urlsDB     *mongo.Collection
	app        *fiber.App
)

func clientsInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(mongoDBURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	urlsDB = client.Database("shawty").Collection("urls")
}

func serverInit() {
	engine := html.New("./web/template", ".html")
	app = fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Static("/", "./web/static")
}

func main() {
	clientsInit()
	serverInit()
	app.Get("/", handlers.HomePage)
	app.Post("/", handlers.AddURL)
	app.Get("/:shortURL", handlers.LongURLRedirect)
	log.Fatalln(app.Listen(serverURL[7:]))
}
