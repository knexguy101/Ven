package user

import (
	"Ven/requests"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BuyItemPayload struct {
	Amount int `json:"amount"`
}

type BuyItemResponse struct {
	Message string `json:"message"`
}

func BuyItem(itemId string, amount int, client *requests.HttpClient) (*BuyItemResponse, error) {

	payload, _ := json.Marshal(BuyItemPayload {
		Amount: amount,
	})

	req, _ := http.NewRequestWithContext(client.Context, "POST", fmt.Sprintf("https://api.legendsofvenari.com/shop/%s", itemId), bytes.NewBuffer(payload))
	req.Header = map[string][]string {
		"Accept": {"application/json, text/plain, */*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"Content-Type": {"application/json"},
		"X-XSRF-TOKEN": {client.GetXSRF()},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
		"Host": {"api.legendsofvenari.com"},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unknown response on csrf: %d", res.StatusCode)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var userRes BuyItemResponse
	err = json.Unmarshal(resBytes, &userRes)
	if err != nil {
		return nil, err
	}

	return &userRes, nil
}

type GetStockResponse []struct {
	ID   string `json:"id"`
	Item struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Tier        string `json:"tier"`
		Assets      struct {
			Image string `json:"image"`
		} `json:"assets"`
	} `json:"item"`
	Type  string `json:"type"`
	Price struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	} `json:"price"`
	Stock int `json:"stock"`
}

func GetStock(city string, client *requests.HttpClient) (*GetStockResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "GET", fmt.Sprintf("https://api.legendsofvenari.com/shop?area=%s", city), nil)
	req.Header = map[string][]string {
		"Accept": {"application/json, text/plain, */*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"X-XSRF-TOKEN": {client.GetXSRF()},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
		"Host": {"api.legendsofvenari.com"},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unknown response on csrf: %d", res.StatusCode)
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var userRes GetStockResponse
	err = json.Unmarshal(resBytes, &userRes)
	if err != nil {
		return nil, err
	}

	return &userRes, nil
}
