package worlds

import (
	"Ven/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WorldsResponse []struct {
	Costs struct {
		Travel struct {
			Type   string `json:"type"`
			Amount int    `json:"amount"`
		} `json:"travel"`
	} `json:"costs"`
	Assets struct {
		Video string `json:"video"`
		Image string `json:"image"`
	} `json:"assets"`
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Options     struct {
		ColorAccent string `json:"colorAccent"`
	} `json:"options"`
	V         int       `json:"__v"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Areas     []struct {
		Costs struct {
			Travel struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"travel"`
			Expedition struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"expedition"`
		} `json:"costs"`
		UnlockRequirement struct {
			Type  string `json:"type"`
			Level int    `json:"level"`
			World string `json:"world"`
		} `json:"unlockRequirement"`
		ID          string    `json:"_id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Image       string    `json:"image"`
		World       string    `json:"world"`
		V           int       `json:"__v"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
		Venari      []string  `json:"venari"`
		ShopStock   []struct {
			ID   string `json:"id"`
			Item struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Type   string `json:"type"`
				Tier   string `json:"tier"`
				Assets struct {
					Image string `json:"image"`
				} `json:"assets"`
			} `json:"item"`
			Type  string `json:"type"`
			Price struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"price"`
			Stock int `json:"stock"`
		} `json:"shopStock"`
	} `json:"areas"`
	UnlockRequirement struct {
		Type  string `json:"type"`
		Level int    `json:"level"`
		World string `json:"world,omitempty"`
	} `json:"unlockRequirement,omitempty"`
}

func GetWorlds(client *requests.HttpClient) (*WorldsResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "GET", "https://api.legendsofvenari.com/worlds", nil)
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

	var worldRes WorldsResponse
	err = json.Unmarshal(resBytes, &worldRes)
	if err != nil {
		return nil, err
	}

	return &worldRes, nil
}
