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
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	authRepo := auth.NewRepository(cfg.Auth, cfg.DB)
	opt := option.WithCredentialsFile("firebase-key.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Firebase initialization error: %v", err)
	}

	app := fiber.New(fiber.Config{
		ReadBufferSize: 1024 * 10,
	})

	app.Use(cors.New(cors.Config{
    		AllowOrigins:     "http://localhost:3000, https://trader.pro-fe.abhinavganeshan.in",
    		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
    		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
    		AllowCredentials: true,
	}))
	

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "hello"})
	})

	app.Use(middlewares.FirebaseAuthMiddleware(firebaseApp))

	routes.RegisterAuthRoutes(app, authRepo)

	SimulationHandler := api.NewSimulationHandler(cfg.DB.DB)
	routes.MountSimulationRoutes(app, SimulationHandler)

	log.Println("Server running on http://localhost:8080")

	log.Fatal(app.Listen(":8080"))
}
