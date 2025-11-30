package api

import (
	"api-gateway/config"
	"api-gateway/internal/api/controllers"
	"api-gateway/internal/api/routes"
	"api-gateway/internal/client"
	"api-gateway/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize Clients
	walletClient := client.NewWalletClient(cfg.WalletURL)
	giftClient := client.NewGiftClient(cfg.GiftURL)

	// Initialize Service
	gatewayService := service.NewGatewayService(walletClient, giftClient)

	// Initialize Controller
	gatewayController := controllers.NewGatewayController(gatewayService)

	// Fiber App
	app := fiber.New()

	// Register Routes
	routes.RegisterRoutes(app, gatewayController)

	log.Println("🌐 API Gateway running on port", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
