package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"

	"github.com/lordlabuckdas/shawty/config"
	"github.com/lordlabuckdas/shawty/storage"
	"github.com/lordlabuckdas/shawty/utils"
)

func HomePage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func AddURL(c *fiber.Ctx) error {
	longURL, err := url.Parse(c.FormValue("longURL"))
	if len(longURL.String()) == 0 {
		return HomePage(c)
	}
	if err != nil {
		utils.ErrorLogger.Println(err)
		return utils.HandleError(c, &utils.URLError{
			Code: 400,
			Msg:  "Invalid URL",
		})
	}
	utils.InfoLogger.Println("URL received: " + longURL.String())
	if len(longURL.Scheme) == 0 {
		longURL.Scheme = "http"
	}
	shortURL, err := utils.ShortenURL(longURL.String())
	if err != nil {
		utils.ErrorLogger.Println(err)
		return utils.HandleError(c, err)
	}
	utils.InfoLogger.Println("Shortened URL: " + shortURL)
	err = storage.Insert(shortURL, longURL.String(), c)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return utils.HandleError(c, err)
	}
	utils.InfoLogger.Println("Inserted URL to DB")
	return c.Render("index", fiber.Map{
		"longURL":  longURL,
		"shortURL": config.ServerURL + "/" + shortURL,
	})
}

func LongURLRedirect(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")
	utils.InfoLogger.Println("URL received: " + shortURL)
	longURL, err := storage.Find(shortURL, c)
	if err != nil {
		utils.ErrorLogger.Println(err)
		return utils.HandleError(c, err)
	}
	utils.InfoLogger.Println("Long URL found: " + longURL)
	return c.Status(302).Redirect(longURL)
}
