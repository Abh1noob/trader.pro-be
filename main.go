package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/Abh1noob/trader.pro-be/api"
	"github.com/Abh1noob/trader.pro-be/config"
	"github.com/Abh1noob/trader.pro-be/internal/auth"
	"github.com/Abh1noob/trader.pro-be/middlewares"
	"github.com/Abh1noob/trader.pro-be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/api/option"
)

func main() {
	// Load Config
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Firebase Auth
	authRepo := auth.NewRepository(cfg.Auth, cfg.DB)
	opt := option.WithCredentialsFile("firebase-key.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v", err)
	}

	// Initialize Fiber App
	app := fiber.New(fiber.Config{
		ReadBufferSize: 1024 * 10,
	})

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000,https://yourproductiondomain.com",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Use Firebase Middleware
	app.Use(middlewares.FirebaseAuthMiddleware(firebaseApp))

	// Initialize Handlers

	// Register Routes
	routes.RegisterAuthRoutes(app, authRepo)

	SimulationHandler := api.NewSimulationHandler(cfg.DB.DB)
	routes.MountSimulationRoutes(app, SimulationHandler)

	// Public Route Example
	app.Get("/public/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a public route"})
	})

	// Start Server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
