package repository

import (
	"wallet-service/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(txn *models.Transaction) error {
	return r.db.Create(txn).Error
}

func (r *TransactionRepository) GetByWalletID(walletID string) ([]models.Transaction, error) {
	var txns []models.Transaction
	err := r.db.Where("wallet_id = ?", walletID).Find(&txns).Error
	return txns, err
}
