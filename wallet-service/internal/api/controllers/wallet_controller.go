package controllers

import (
	"net/http"
	"wallet-service/internal/service"

	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	WalletService *service.WalletService
}

func NewWalletController(service *service.WalletService) *WalletController {
	return &WalletController{WalletService: service}
}

// POST /wallet/create
func (wc *WalletController) CreateWallet(ctx *fiber.Ctx) error {
	var body struct {
		Phone string `json:"phone"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	wallet, err := wc.WalletService.CreateWallet(body.Phone)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(wallet)
}

// GET /wallet/:phone
func (wc *WalletController) GetWallet(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")
	wallet, err := wc.WalletService.GetWalletByPhone(phone)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if wallet == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "wallet not found"})
	}
	return ctx.JSON(wallet)
}

// POST /wallet/add-balance
func (wc *WalletController) AddBalance(ctx *fiber.Ctx) error {
	var body struct {
		Phone     string `json:"phone"`
		Amount    int64  `json:"amount"`
		Reference string `json:"reference"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	wallet, err := wc.WalletService.AddBalance(body.Phone, body.Amount, body.Reference)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(wallet)
}

// GET /wallet/list
func (wc *WalletController) ListWallets(ctx *fiber.Ctx) error {
	wallets, err := wc.WalletService.ListWallets()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(wallets)
}

// GET /wallet/:phone/transactions
func (wc *WalletController) GetTransactions(ctx *fiber.Ctx) error {
	phone := ctx.Params("phone")
	txns, err := wc.WalletService.GetTransactions(phone)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(txns)
}
