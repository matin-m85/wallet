package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WalletClient struct {
	BaseURL string
	Client  *http.Client
}

func NewWalletClient(baseURL string) *WalletClient {
	return &WalletClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 5 * time.Second},
	}
}

type Wallet struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Balance int64  `json:"balance"`
}

type Transaction struct {
	ID        string `json:"id"`
	WalletID  string `json:"wallet_id"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
	CreatedAt string `json:"created_at"`
}

type CreateWalletRequest struct {
	Phone string `json:"phone"`
}

type AddBalanceRequest struct {
	Phone     string `json:"phone"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
}

// Create Wallet
func (wc *WalletClient) CreateWallet(phone string) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet", wc.BaseURL)
	body := CreateWalletRequest{Phone: phone}
	jsonBody, _ := json.Marshal(body)

	resp, err := wc.Client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wallet Wallet
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Add Balance
func (wc *WalletClient) AddBalance(phone string, amount int64, reference string) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet/add", wc.BaseURL)

	body := AddBalanceRequest{Phone: phone, Amount: amount, Reference: reference}
	jsonBody, _ := json.Marshal(body)

	resp, err := wc.Client.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wallet Wallet
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Get Wallet
func (wc *WalletClient) GetWallet(phone string) (*Wallet, error) {
	url := fmt.Sprintf("%s/wallet/%s", wc.BaseURL, phone)
	resp, err := wc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wallet Wallet
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}

// List Wallets
func (wc *WalletClient) ListWallets() ([]Wallet, error) {
	url := fmt.Sprintf("%s/wallet/list", wc.BaseURL)
	resp, err := wc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var wallets []Wallet
	if err := json.NewDecoder(resp.Body).Decode(&wallets); err != nil {
		return nil, err
	}
	return wallets, nil
}

// Get Transactions
func (wc *WalletClient) GetTransactions(phone string) ([]Transaction, error) {
	url := fmt.Sprintf("%s/wallet/%s/transactions", wc.BaseURL, phone)

	resp, err := wc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var txs []Transaction
	if err := json.NewDecoder(resp.Body).Decode(&txs); err != nil {
		return nil, err
	}
	return txs, nil
}
