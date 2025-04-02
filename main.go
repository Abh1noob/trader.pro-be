package main

import (
	"log"

	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/Abh1noob/trader.pro-be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	authRepo := auth.NewRepository(cfg.Auth, cfg.DB)

	app := fiber.New(fiber.Config{
		ReadBufferSize: 1024 * 10, // Increase buffer size (10 KB)
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,https://yourproductiondomain.com",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.RegisterAuthRoutes(app, authRepo)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
