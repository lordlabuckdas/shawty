package storage

import (
	"github.com/gofiber/fiber/v2"

	// "github.com/go-redis/redis"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lordlabuckdas/shawty/utils"
)

func Insert(shortURL string, longURL string, c *fiber.Ctx) error {
	urlItem := &URL{
		LongURL: longURL,
		ShortURL: shortURL,
	}
	utils.InfoLogger.Println("Inserting " + shortURL + " and " + longURL)
	bsonURL, err := bson.Marshal(urlItem)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return &utils.URLError{
			Code: 500,
			Msg: "Error marshalling to BSON",
		}
	}
	utils.InfoLogger.Println("Successfully marshalled")
	_, err = urlsDB.InsertOne(c.Context(), bsonURL)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return &utils.URLError{
			Code: 500,
			Msg: "Error marshalling to BSON",
		}
	}
	utils.InfoLogger.Println("Successfully inserted into DB")
	return nil
}

func Find(shortURL string, c *fiber.Ctx) (string, error) {
	utils.InfoLogger.Println("Attempting to find " + shortURL + " in DB")
	urlItem := &URL{}
	err := urlsDB.FindOne(c.Context(), bson.D{primitive.E{
		Key:   "short_url",
		Value: shortURL,
	}}).Decode(&urlItem)
	if err != nil {
		utils.ErrorLogger.Println(err)
		if err == mongo.ErrNoDocuments {
			return "", &utils.URLError{
				Code: 404,
				Msg: "Page not found!",
			}
		}
		return "", &utils.URLError{
			Code: 500,
			Msg: "Internal Error",
		}
	}
	utils.InfoLogger.Println("URL found in DB")
	return urlItem.LongURL, nil
}
