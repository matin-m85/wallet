package service

import (
	"errors"
	"wallet-service/internal/models"
	"wallet-service/internal/repository"
)

type WalletService struct {
	walletRepo      *repository.WalletRepository
	transactionRepo *repository.TransactionRepository
}

func NewWalletService(walletRepo *repository.WalletRepository, transactionRepo *repository.TransactionRepository) *WalletService {
	return &WalletService{
		walletRepo:      walletRepo,
		transactionRepo: transactionRepo,
	}
}

func (WalletService *WalletService) CreateWallet(phone string) (*models.Wallet, error) {
	existing, err := WalletService.walletRepo.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("wallet with this phone already exists")
	}

	wallet := &models.Wallet{
		Phone: phone,
	}

	if err := WalletService.walletRepo.Create(wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (WalletService *WalletService) GetWalletByPhone(phone string) (*models.Wallet, error) {
	return WalletService.walletRepo.GetByPhone(phone)
}

func (WalletService *WalletService) AddBalance(phone string, amount int64, reference string) (*models.Wallet, error) {
	wallet, err := WalletService.walletRepo.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}

	if err := WalletService.walletRepo.UpdateBalance(wallet, amount); err != nil {
		return nil, err
	}

	txn := &models.Transaction{
		WalletID:  wallet.ID,
		Amount:    amount,
		Reference: reference,
	}
	if err := WalletService.transactionRepo.Create(txn); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (WalletService *WalletService) GetTransactions(phone string) ([]models.Transaction, error) {
	wallet, err := WalletService.walletRepo.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	if wallet == nil {
		return nil, errors.New("wallet not found")
	}
	return WalletService.transactionRepo.GetByWalletID(wallet.ID.String())
}

func (s *WalletService) ListWallets() ([]models.Wallet, error) {
	return s.walletRepo.ListAll()
}
