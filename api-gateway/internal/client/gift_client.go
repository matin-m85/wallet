package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GiftClient struct {
	BaseURL string
	Client  *http.Client
}

func NewGiftClient(baseURL string) *GiftClient {
	return &GiftClient{
		BaseURL: baseURL,
		Client:  &http.Client{Timeout: 5 * time.Second},
	}
}

type CreateGroupRequest struct {
	Name   string `json:"name"`
	Amount int64  `json:"amount_due"`
	Count  int    `json:"count"`
}

type CreateGroupResponse struct {
	GroupID   string   `json:"group_id"`
	Name      string   `json:"name"`
	AmountDue int64    `json:"amount_due"`
	Count     int      `json:"count"`
	Codes     []string `json:"codes"`
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

// Create Gift Group
func (gc *GiftClient) CreateGroup(req CreateGroupRequest) (*CreateGroupResponse, error) {
	url := fmt.Sprintf("%s/gift/group/create", gc.BaseURL)
	body, _ := json.Marshal(req)

	resp, err := gc.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out CreateGroupResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Use Gift Code
func (gc *GiftClient) UseCode(req UseCodeRequest) (*UseCodeResponse, error) {
	url := fmt.Sprintf("%s/gift/use", gc.BaseURL)
	body, _ := json.Marshal(req)

	resp, err := gc.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out UseCodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Group Stats
func (gc *GiftClient) GroupStats(groupID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gift/group/%s/stats", gc.BaseURL, groupID)
	resp, err := gc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// Group Users
func (gc *GiftClient) GroupUsers(groupID string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gift/group/%s/users", gc.BaseURL, groupID)
	resp, err := gc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// Card Info
func (gc *GiftClient) CardInfo(code string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gift/card/%s", gc.BaseURL, code)
	resp, err := gc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// List All Groups
func (gc *GiftClient) ListGroups() ([]map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gift/group/list", gc.BaseURL)
	resp, err := gc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// List Codes by Group
func (gc *GiftClient) ListGroupCodes(groupID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/gift/group/%s/codes", gc.BaseURL, groupID)
	resp, err := gc.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}
