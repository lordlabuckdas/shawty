package storage

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/lordlabuckdas/shawty/config"
	"github.com/lordlabuckdas/shawty/utils"
)

type URL struct {
	// Clicks uint64 `json:"clicks" bson:"clicks"`
	// CreationDate string `json:"creation_date" bson:"creation_date"`
	LongURL  string `json:"long_url" bson:"long_url"`
	ShortURL string `json:"short_url" bson:"short_url"`
}

var (
	urlsDB *mongo.Collection
)

func ClientsInit() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(config.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return &utils.URLError{
			Code: 500,
			Msg:  "Cannot establish connection to MongoDB",
		}
	}
	urlsDB = client.Database(config.DatabaseName).Collection(config.CollectionName)
	return nil
}
