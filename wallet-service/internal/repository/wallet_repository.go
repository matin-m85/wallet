package repository

import (
	"errors"
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) Create(wallet *models.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepository) GetByID(id string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.First(&wallet, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) GetByPhone(phone string) (*models.Wallet, error) {
	var wallet models.Wallet
	if err := r.db.First(&wallet, "phone = ?", phone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func (r *WalletRepository) UpdateBalance(wallet *models.Wallet, amount int64) error {
	wallet.Balance += amount
	return r.db.Save(wallet).Error
}

func (r *WalletRepository) ListAll() ([]models.Wallet, error) {
	var wallets []models.Wallet
	if err := r.db.Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}
