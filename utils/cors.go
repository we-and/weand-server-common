package utils

import (
	fib "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	helmet_ "github.com/gofiber/helmet/v2"
)

func SetHeaders(app *fib.App) {
	app.Use(helmet_.New())
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowOrigins: "*", //"https://www.pool.ca",
		AllowHeaders: "*",
	}))
}
