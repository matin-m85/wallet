package controllers

import (
	"api-gateway/internal/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GatewayController struct {
	service *service.GatewayService
}

func NewGatewayController(s *service.GatewayService) *GatewayController {
	return &GatewayController{service: s}
}

// ---------------- Wallet ----------------
func (c *GatewayController) CreateWallet(ctx *fiber.Ctx) error {
	var body struct {
		Phone string `json:"phone"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	wallet, err := c.service.CreateWallet(body.Phone)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(wallet)
}

func (c *GatewayController) AddBalance(ctx *fiber.Ctx) error {
	var body struct {
		Phone     string `json:"phone"`
		Amount    int64  `json:"amount"`
		Reference string `json:"reference"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	wallet, err := c.service.AddBalance(body.Phone, body.Amount, body.Reference)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(wallet)
}

func (c *GatewayController) GetWallet(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")
	wallet, err := c.service.GetWallet(phone)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(wallet)
}

func (c *GatewayController) ListWallets(ctx *fiber.Ctx) error {
	wallets, err := c.service.ListWallets()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(wallets)
}

func (c *GatewayController) GetTransactions(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")
	txs, err := c.service.GetTransactions(phone)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(txs)
}

// ---------------- Gift ----------------
func (c *GatewayController) CreateGiftGroup(ctx *fiber.Ctx) error {
	var body struct {
		Name   string `json:"name"`
		Amount int64  `json:"amount_due"`
		Count  int    `json:"count"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	res, err := c.service.CreateGiftGroup(body.Name, body.Amount, body.Count)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (c *GatewayController) UseGiftCode(ctx *fiber.Ctx) error {
	var body struct {
		Code  string `json:"code"`
		Phone string `json:"phone"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}
	res, wallet, err := c.service.UseGiftCode(body.Code, body.Phone)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"wallet": wallet,
		"gift":   res,
	})
}

func (c *GatewayController) GiftGroupStats(ctx *fiber.Ctx) error {
	groupID := ctx.Params("group_id")
	res, err := c.service.GiftGroupStats(groupID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (c *GatewayController) GiftGroupUsers(ctx *fiber.Ctx) error {
	groupID := ctx.Params("group_id")
	res, err := c.service.GiftGroupUsers(groupID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (c *GatewayController) GiftCardInfo(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	res, err := c.service.GiftCardInfo(code)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (c *GatewayController) GiftGroupsList(ctx *fiber.Ctx) error {
	res, err := c.service.GiftGroupsList()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}

func (c *GatewayController) GiftGroupCodes(ctx *fiber.Ctx) error {
	groupID := ctx.Params("group_id")
	res, err := c.service.GiftGroupCodes(groupID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(res)
}
