package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/notes-api/url"
)

func main() {
	app := fiber.New()

	url.Web(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}