package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/notes-api/url"
	"os"
)

func main() {
	app := fiber.New()

	url.Web(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}