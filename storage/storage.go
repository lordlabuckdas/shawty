package storage

import (
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lordlabuckdas/shawty/utils"
)

type URL struct {
	// Clicks uint64 `json:"clicks" bson:"clicks"`
	// CreationDate string `json:"creation_date" bson:"creation_date"`
	LongURL  string `json:"long_url" bson:"long_url"`
	ShortURL string `json:"short_url" bson:"short_url"`
}

func Insert(shortURL string, longURL string, c *fiber.Ctx) error {
	urlItem := &URL{
		LongURL: longURL,
		ShortURL: shortURL,
	}
	bsonURL, err := bson.Marshal(urlItem)
	if err != nil {
		return &utils.URLError{500, "Error marshalling to BSON"}
	}
	_, err = urlsDB.InsertOne(c.Context(), bsonURL)
	if err != nil {
		return &utils.URLError{500, "Error marshalling to BSON"}
	}
	return nil
}

func Find(shortURL string) (string, error) {
	// err := urlsDB.FindOne(c.Context(), bson.D{primitive.E{
	// 	Key:   "short_url",
	// 	Value: shortURL,
	// }}).Decode(&url)
	// // log.Println("Executed Search")
	// if err != nil {
	// 	if err == mongo.ErrNoDocuments {
	// 		return c.Status(404).Render("home", fiber.Map{
	// 			"error": "404 - Page Not Found",
	// 		})
	// 	} else {
	// 		return c.Status(400).Render("home", fiber.Map{
	// 			"error": "400 - Bad Request",
	// 		})
	// 	}
	// }
	longURL := "http://example.com"
	return longURL, nil
}
