package main

import (
	"log"

	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pong", "status": 200})
	})

	log.Println("Starting server on :80...")
	if err := app.Listen(":80"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
