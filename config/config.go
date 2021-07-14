package config

import "os"

var (
	ServerURL  string = "http://" + os.Getenv("SHAWTY_HOST") + ":" + os.Getenv("SHAWTY_PORT")
	MongoDBURI string = os.Getenv("SHAWTY_MONGODB_URI")
)
