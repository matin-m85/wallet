package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gift-service/internal/models"
	"gift-service/internal/repository"
	"gift-service/util"
)

var (
	ErrCodeNotFound    = errors.New("code not found")
	ErrCodeAlreadyUsed = errors.New("code already used")
)

type Wallet struct {
	ID      string `json:"id"`
	Phone   string `json:"phone"`
	Balance int64  `json:"balance"`
}

type WalletAddRequest struct {
	Phone     string `json:"phone"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
}

type UseCodeRequest struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

type UseCodeResponse struct {
	Code   string `json:"code"`
	Amount int64  `json:"amount"`
	UsedBy string `json:"used_by"`
	UsedAt string `json:"used_at"`
}

type GiftService struct {
	groupRepo *repository.GiftGroupRepo
	cardRepo  *repository.GiftCardRepo
	walletURL string
}

func NewGiftService(gr *repository.GiftGroupRepo, cr *repository.GiftCardRepo, walletURL string) *GiftService {
	return &GiftService{
		groupRepo: gr,
		cardRepo:  cr,
		walletURL: walletURL,
	}
}

func (s *GiftService) CreateGroupAndCards(name string, amount int64, count int) (*models.GiftCardGroup, []models.GiftCard, error) {
	group := &models.GiftCardGroup{
		Name:      name,
		AmountDue: amount,
		CreatedAt: time.Now().UTC(),
	}
	if err := s.groupRepo.Create(group); err != nil {
		return nil, nil, err
	}

	prefix := util.GeneratePrefix(name)

	cards := make([]*models.GiftCard, 0, count)
	for i := 0; i < count; i++ {
		code := util.GenerateCode(prefix)
		card := &models.GiftCard{
			GroupID:   group.ID,
			Code:      code,
			Amount:    amount,
			Used:      false,
			CreatedAt: time.Now().UTC(),
		}
		cards = append(cards, card)
	}

	if err := s.cardRepo.CreateMany(cards); err != nil {
		return group, nil, err
	}

	out := make([]models.GiftCard, len(cards))
	for i, v := range cards {
		out[i] = *v
	}
	return group, out, nil
}

func (s *GiftService) UseGiftCode(req UseCodeRequest) (*UseCodeResponse, *Wallet, error) {
	updatedCard, err := s.cardRepo.UseCardAtomic(req.Code, req.Phone)
	if err != nil {
		return nil, nil, err
	}
	if updatedCard == nil {
		c, err := s.cardRepo.GetByCode(req.Code)
		if err != nil {
			return nil, nil, err
		}
		if c == nil {
			return nil, nil, ErrCodeNotFound
		}
		return nil, nil, ErrCodeAlreadyUsed
	}

	group, _ := s.groupRepo.GetByID(updatedCard.GroupID.String())
	amount := updatedCard.Amount
	if group != nil {
		amount = group.AmountDue
	}

	res := &UseCodeResponse{
		Code:   updatedCard.Code,
		Amount: amount,
		UsedBy: req.Phone,
		UsedAt: updatedCard.UsedAt.UTC().Format(time.RFC3339),
	}

	walletReq := WalletAddRequest{
		Phone:     req.Phone,
		Amount:    amount,
		Reference: "gift_code:" + req.Code,
	}
	body, _ := json.Marshal(walletReq)
	resp, err := http.Post(fmt.Sprintf("%s/wallet/add", s.walletURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var wallet Wallet
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		return nil, nil, err
	}

	return res, &wallet, nil
}

func (s *GiftService) GetGroupStats(groupID string) (map[string]interface{}, error) {
	group, err := s.groupRepo.GetByID(groupID)
	if err != nil {
		return nil, err
	}
	if group == nil {
		return nil, nil
	}
	total, used, err := s.groupRepo.GetStats(groupID)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"group_id":        group.ID,
		"name":            group.Name,
		"amount_due":      group.AmountDue,
		"total_cards":     total,
		"used_cards":      used,
		"remaining_cards": total - used,
	}, nil
}

func (s *GiftService) ListGroupUsers(groupID string) ([]map[string]interface{}, error) {
	usedCards, err := s.cardRepo.ListUsedByGroup(groupID)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]interface{}, 0, len(usedCards))
	for _, c := range usedCards {
		out = append(out, map[string]interface{}{
			"phone":   c.UserPhone,
			"code":    c.Code,
			"used_at": c.UsedAt,
		})
	}
	return out, nil
}

func (s *GiftService) GetCardInfo(code string) (map[string]interface{}, error) {
	c, err := s.cardRepo.GetByCode(code)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, nil
	}
	group, _ := s.groupRepo.GetByID(c.GroupID.String())
	res := map[string]interface{}{
		"code":       c.Code,
		"group_name": "",
		"amount_due": int64(0),
		"used":       c.Used,
		"user_phone": c.UserPhone,
		"used_at":    c.UsedAt,
	}
	if group != nil {
		res["group_name"] = group.Name
		res["amount_due"] = group.AmountDue
	}
	return res, nil
}

func (s *GiftService) ListGroupsSummary() ([]map[string]interface{}, error) {
	groups, err := s.groupRepo.ListAll()
	if err != nil {
		return nil, err
	}
	out := make([]map[string]interface{}, 0, len(groups))
	for _, g := range groups {
		total, used, err := s.groupRepo.GetStats(g.ID.String())
		if err != nil {
			return nil, err
		}
		out = append(out, map[string]interface{}{
			"group_id":        g.ID,
			"name":            g.Name,
			"amount_due":      g.AmountDue,
			"total_cards":     total,
			"used_cards":      used,
			"remaining_cards": total - used,
		})
	}
	return out, nil
}

func (s *GiftService) ListGroupCodes(groupID string) (map[string]interface{}, error) {
	cards, err := s.cardRepo.ListByGroup(groupID)
	if err != nil {
		return nil, err
	}
	usedCodes := []string{}
	unusedCodes := []string{}
	for _, c := range cards {
		if c.Used {
			usedCodes = append(usedCodes, c.Code)
		} else {
			unusedCodes = append(unusedCodes, c.Code)
		}
	}
	return map[string]interface{}{
		"used_count":   len(usedCodes),
		"unused_count": len(unusedCodes),
		"used_codes":   usedCodes,
		"unused_codes": unusedCodes,
	}, nil
}
