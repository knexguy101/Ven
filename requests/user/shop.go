package user

import (
	"VenariBot/requests"
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
