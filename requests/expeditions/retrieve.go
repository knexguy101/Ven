package expeditions

import (
	"Ven/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ExpeditionResponse []struct {
	ID       string `json:"_id"`
	Location struct {
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
		ID          string `json:"_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
		World       struct {
			UnlockRequirement struct {
				Type  string `json:"type"`
				Level int    `json:"level"`
			} `json:"unlockRequirement"`
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
		} `json:"world"`
		Shop []struct {
			Price struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"price"`
			Item  string `json:"item"`
			Model string `json:"model"`
			Stock int    `json:"stock"`
			ID    string `json:"_id"`
		} `json:"shop"`
		V         int       `json:"__v"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		ShopStock []struct {
			ID    string `json:"id"`
			Item  string `json:"item"`
			Type  string `json:"type"`
			Price struct {
				Type   string `json:"type"`
				Amount int    `json:"amount"`
			} `json:"price"`
			Stock int `json:"stock"`
		} `json:"shopStock"`
	} `json:"location"`
	Bait struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Tier        string `json:"tier"`
		Assets      struct {
			Image string `json:"image"`
		} `json:"assets"`
	} `json:"bait"`
	User       string        `json:"user"`
	SpawnRolls int           `json:"spawnRolls"`
	Inventory  []interface{} `json:"inventory"`
	Spawns     []struct {
		Venari struct {
			ID           string `json:"_id"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Tier         string `json:"tier"`
			ModelOptions struct {
				Bloom int     `json:"bloom"`
				Scale float64 `json:"scale"`
			} `json:"modelOptions"`
			Bio    string `json:"bio"`
			Assets struct {
				Image       string `json:"image"`
				Avatar      string `json:"avatar"`
				Video       string `json:"video"`
				LargeVideo  string `json:"largeVideo"`
				Model       string `json:"model"`
				EmissionMap string `json:"emissionMap"`
				Texture     string `json:"texture"`
			} `json:"assets"`
		} `json:"venari"`
		Hidden bool   `json:"hidden"`
		ID     string `json:"_id"`
	} `json:"spawns"`
	Rewards          []interface{} `json:"rewards"`
	BaitingStartedAt time.Time     `json:"baitingStartedAt"`
	V                int           `json:"__v"`
	BaitDuration     int           `json:"baitDuration"`
	CaptureChance    string        `json:"captureChance"`
}

func GetExpeditions(area string, client *requests.HttpClient) (*ExpeditionResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "GET", fmt.Sprintf("https://api.legendsofvenari.com/expeditions?area=%s&active=true", area), nil)
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

	var expRes ExpeditionResponse
	err = json.Unmarshal(resBytes, &expRes)
	if err != nil {
		return nil, err
	}

	return &expRes, nil
}