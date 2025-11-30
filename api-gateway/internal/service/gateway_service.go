package service

import "api-gateway/internal/client"

type GatewayService struct {
	walletClient *client.WalletClient
	giftClient   *client.GiftClient
}

func NewGatewayService(wc *client.WalletClient, gc *client.GiftClient) *GatewayService {
	return &GatewayService{walletClient: wc, giftClient: gc}
}

// ---------------- Wallet ----------------

func (s *GatewayService) CreateWallet(phone string) (*client.Wallet, error) {
	return s.walletClient.CreateWallet(phone)
}

func (s *GatewayService) AddBalance(phone string, amount int64, reference string) (*client.Wallet, error) {
	return s.walletClient.AddBalance(phone, amount, reference)
}

func (s *GatewayService) GetWallet(phone string) (*client.Wallet, error) {
	return s.walletClient.GetWallet(phone)
}

func (s *GatewayService) ListWallets() ([]client.Wallet, error) {
	return s.walletClient.ListWallets()
}

func (s *GatewayService) GetTransactions(phone string) ([]client.Transaction, error) {
	return s.walletClient.GetTransactions(phone)
}

// ---------------- Gift ----------------

func (s *GatewayService) CreateGiftGroup(name string, amount int64, count int) (*client.CreateGroupResponse, error) {
	req := client.CreateGroupRequest{Name: name, Amount: amount, Count: count}
	return s.giftClient.CreateGroup(req)
}

func (s *GatewayService) UseGiftCode(code, phone string) (*client.UseCodeResponse, *client.Wallet, error) {
	res, err := s.giftClient.UseCode(client.UseCodeRequest{Code: code, Phone: phone})
	if err != nil {
		return nil, nil, err
	}

	wallet, err := s.walletClient.AddBalance(phone, res.Amount, "gift_code:"+code)
	if err != nil {
		return nil, nil, err
	}

	return res, wallet, nil
}

func (s *GatewayService) GiftGroupStats(groupID string) (map[string]interface{}, error) {
	return s.giftClient.GroupStats(groupID)
}

func (s *GatewayService) GiftGroupUsers(groupID string) ([]map[string]interface{}, error) {
	return s.giftClient.GroupUsers(groupID)
}

func (s *GatewayService) GiftCardInfo(code string) (map[string]interface{}, error) {
	return s.giftClient.CardInfo(code)
}

func (s *GatewayService) GiftGroupsList() ([]map[string]interface{}, error) {
	return s.giftClient.ListGroups()
}

func (s *GatewayService) GiftGroupCodes(groupID string) (map[string]interface{}, error) {
	return s.giftClient.ListGroupCodes(groupID)
}
