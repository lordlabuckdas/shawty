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
	if err != nil {
		return utils.HandleError(c, &utils.URLError{
			Code: 400,
			Msg: "Invalid URL",
		})
	}
	if len(longURL.Scheme) == 0 {
		longURL.Scheme = "http"
	}
	shortURL, err := utils.ShortenURL(longURL.String())
	if err != nil {
		return utils.HandleError(c, err)
	}
	err = storage.Insert(shortURL, longURL, c)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.Render("index", fiber.Map{
		"longURL":  longURL,
		"shortURL": config.ServerURL + "/" + shortURL,
	})
}

func LongURLRedirect(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")
	longURL, err := storage.Find(shortURL)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return c.Status(302).Redirect(longURL)
}
