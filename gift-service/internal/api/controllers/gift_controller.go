package controllers

import (
	"gift-service/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GiftController struct {
	GiftService *service.GiftService
}

func NewGiftController(s *service.GiftService) *GiftController {
	return &GiftController{GiftService: s}
}

func (gc *GiftController) CreateGroup(ctx *fiber.Ctx) error {
	var body struct {
		Name   string `json:"name"`
		Amount int64  `json:"amount_due"`
		Count  int    `json:"count"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	if body.Count <= 0 {
		body.Count = 1
	}
	group, cards, err := gc.GiftService.CreateGroupAndCards(body.Name, body.Amount, body.Count)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	codes := make([]string, 0, len(cards))
	for _, card := range cards {
		codes = append(codes, card.Code)
	}
	return ctx.JSON(fiber.Map{
		"group_id":   group.ID,
		"name":       group.Name,
		"amount_due": group.AmountDue,
		"count":      len(cards),
		"codes":      codes,
	})
}

func (gc *GiftController) UseCode(ctx *fiber.Ctx) error {
	var body struct {
		Code  string `json:"code"`
		Phone string `json:"phone"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	res, wallet, err := gc.GiftService.UseGiftCode(service.UseCodeRequest{
		Code:  body.Code,
		Phone: body.Phone,
	})
	if err != nil {
		switch err {
		case service.ErrCodeAlreadyUsed:
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "code already used"})
		case service.ErrCodeNotFound:
			return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "code not found"})
		default:
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	}
	return ctx.JSON(fiber.Map{
		"gift":   res,
		"wallet": wallet,
	})
}

func (gc *GiftController) GroupStats(ctx *fiber.Ctx) error {
	groupID := ctx.Params("group_id")
	stat, err := gc.GiftService.GetGroupStats(groupID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if stat == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "group not found"})
	}
	return ctx.JSON(stat)
}

func (gc *GiftController) GroupUsers(ctx *fiber.Ctx) error {
	groupID := ctx.Params("group_id")
	list, err := gc.GiftService.ListGroupUsers(groupID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(list)
}

func (gc *GiftController) CardInfo(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	info, err := gc.GiftService.GetCardInfo(code)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if info == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "card not found"})
	}
	return ctx.JSON(info)
}

func (gc *GiftController) GroupList(ctx *fiber.Ctx) error {
	list, err := gc.GiftService.ListGroupsSummary()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(list)
}

func (gc *GiftController) ListGroupCodes(ctx *fiber.Ctx) error {
	groupID := ctx.Params("id")
	res, err := gc.GiftService.ListGroupCodes(groupID)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}
