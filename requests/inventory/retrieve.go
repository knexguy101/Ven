package inventory

import (
	"Ven/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type InventoryResponse []struct {
	ID   string `json:"_id"`
	Item struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Tier        string `json:"tier"`
		Assets      struct {
			Image string `json:"image"`
		} `json:"assets"`
		Amount int `json:"amount"`
	} `json:"item,omitempty"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetInventory(client *requests.HttpClient) (*InventoryResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "GET", "https://api.legendsofvenari.com/inventory", nil)
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

	var invenRes InventoryResponse
	err = json.Unmarshal(resBytes, &invenRes)
	if err != nil {
		return nil, err
	}

	return &invenRes, nil
}