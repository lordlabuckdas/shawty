package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/lordlabuckdas/shawty/config"
	"github.com/lordlabuckdas/shawty/handlers"
	"github.com/lordlabuckdas/shawty/storage"
	"github.com/lordlabuckdas/shawty/utils"
)

var (
	app        *fiber.App
)

func main() {
	err := utils.LoggerInit()
	if err != nil {
		log.Fatalln("Error initializing logger")
	}
	err = storage.ClientsInit()
	if err != nil {
		utils.ErrorLogger.Fatalln(err)
	}
	app = handlers.ServerInit()
	app.Get("/", handlers.HomePage)
	app.Post("/", handlers.AddURL)
	app.Get("/:shortURL", handlers.LongURLRedirect)
	utils.InfoLogger.Println("Starting web server")
	utils.ErrorLogger.Fatalln(app.Listen(config.ServerURL[7:]))
}
