package config

import (
	"strings"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://nekowawolf.xyz",
	"https://nekowawolf.github.io",
}

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins[:], ","),
	AllowHeaders:     "Origin,Login,Content-Type",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
}
