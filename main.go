package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekowawolf/notes-api/url"
	"github.com/nekowawolf/notes-api/config" 
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {
	app := fiber.New()
	
	app.Use(cors.New(config.Cors))

	url.Web(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}