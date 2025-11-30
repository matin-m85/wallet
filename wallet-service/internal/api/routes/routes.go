package routes

import (
	"wallet-service/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterWalletRoutes(app *fiber.App, wc *controllers.WalletController) {
	wallet := app.Group("/wallet")

	wallet.Post("/", wc.CreateWallet)
	wallet.Post("/add", wc.AddBalance)

	wallet.Get("/list", wc.ListWallets)
	wallet.Get("/:phone", wc.GetWallet)

	wallet.Get("/:phone/transactions", wc.GetTransactions)

}
