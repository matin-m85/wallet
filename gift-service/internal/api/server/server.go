package server

import (
	"gift-service/config"
	"gift-service/internal/api/controllers"
	"gift-service/internal/api/routes"
	"gift-service/internal/repository"
	"gift-service/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	cfg := config.LoadConfig()

	// Connect to DB
	repository.ConnectDB(cfg)

	// Initialize Repositories
	groupRepo := repository.NewGiftGroupRepo(repository.DB)
	cardRepo := repository.NewGiftCardRepo(repository.DB)

	// Initialize Services (pass wallet url from config)
	giftService := service.NewGiftService(groupRepo, cardRepo, cfg.WalletURL)

	// Initialize Controller
	giftController := controllers.NewGiftController(giftService)

	// Fiber app (use StrictRouting to avoid route collisions)
	app := fiber.New()

	// Register routes
	routes.RegisterGiftRoutes(app, giftController)

	// Read port from ENV (default 8082)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	log.Printf("🌐 Gift Service running on port %s ✅", port)
	log.Fatal(app.Listen(":" + port))
}
