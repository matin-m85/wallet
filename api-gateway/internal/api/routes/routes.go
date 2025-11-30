package routes

import (
	"api-gateway/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, ctrl *controllers.GatewayController) {
	api := app.Group("/api")

	// Wallet
	api.Post("/wallet", ctrl.CreateWallet)
	api.Post("/wallet/add", ctrl.AddBalance)
	api.Get("/wallet/list", ctrl.ListWallets)
	api.Get("/wallet/:phone", ctrl.GetWallet)
	api.Get("/wallet/:phone/transactions", ctrl.GetTransactions)

	// Gift
	api.Post("/group/create", ctrl.CreateGiftGroup)
	api.Post("/use", ctrl.UseGiftCode)
	api.Get("/group/:group_id/stats", ctrl.GiftGroupStats)
	api.Get("/group/:group_id/users", ctrl.GiftGroupUsers)
	api.Get("/card/:code", ctrl.GiftCardInfo)
	api.Get("/group/list", ctrl.GiftGroupsList)
	api.Get("/group/:group_id/codes", ctrl.GiftGroupCodes)
}
