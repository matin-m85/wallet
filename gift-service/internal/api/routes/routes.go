package routes

import (
	"gift-service/internal/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterGiftRoutes(app *fiber.App, gc *controllers.GiftController) {
	gift := app.Group("/gift")

	gift.Post("/group/create", gc.CreateGroup)
	gift.Get("/group/list", gc.GroupList)
	gift.Get("/group/:group_id/stats", gc.GroupStats)
	gift.Get("/group/:group_id/users", gc.GroupUsers)

	gift.Post("/use", gc.UseCode)
	gift.Get("/card/:code", gc.CardInfo)
	gift.Get("/group/:id/codes", gc.ListGroupCodes)

}
