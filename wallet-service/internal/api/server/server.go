package api

import (
	"log"
	"os"
	"wallet-service/config"
	"wallet-service/internal/api/controllers"
	"wallet-service/internal/api/routes"
	"wallet-service/internal/repository"
	"wallet-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to DB
	repository.ConnectDB(cfg)

	// Initialize Repositories
	walletRepo := repository.NewWalletRepository(repository.DB)
	txnRepo := repository.NewTransactionRepository(repository.DB)

	// Initialize Services
	walletService := service.NewWalletService(walletRepo, txnRepo)

	// Initialize Controllers
	walletController := controllers.NewWalletController(walletService)

	// Fiber app (use StrictRouting to avoid route collisions)
	app := fiber.New()

	// Register Routes
	routes.RegisterWalletRoutes(app, walletController)

	// Read port from ENV (default 8081)
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	log.Printf("🌐 Wallet Service running on port %s ✅", port)
	log.Fatal(app.Listen(":" + port))
}
