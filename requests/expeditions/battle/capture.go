package battle

import (
	"Ven/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BattleCatchResponse struct {
	Success bool `json:"success"`
	Rewards []struct {
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	} `json:"rewards"`
	EndedAt time.Time `json:"endedAt"`
}

func BattleCatch(expeditionId string, client *requests.HttpClient) (*BattleCatchResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "POST", fmt.Sprintf("https://api.legendsofvenari.com/expeditions/%s/catch", expeditionId), nil)
	req.Header = map[string][]string{
		"Accept":          {"application/json, text/plain, */*"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"X-XSRF-TOKEN":    {client.GetXSRF()},
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36"},
		"Host":            {"api.legendsofvenari.com"},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unknown response on csrf: %d", res.StatusCode)
	}

	var expRes BattleCatchResponse
	err = json.Unmarshal(resBytes, &expRes)
	if err != nil {
		return nil, err
	}

	return &expRes, nil
}


