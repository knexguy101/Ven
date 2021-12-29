package user

import (
	"VenariBot/requests"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserResponse struct {
	Data struct {
		Experience struct {
			Account struct {
				Amount       int `json:"amount"`
				Level        int `json:"level"`
				Progress     int `json:"progress"`
				NextLevelCap int `json:"nextLevelCap"`
			} `json:"account"`
			World struct {
				Tecta struct {
					Amount       int `json:"amount"`
					Level        int `json:"level"`
					Progress     int `json:"progress"`
					NextLevelCap int `json:"nextLevelCap"`
				} `json:"tecta"`
				Ayena struct {
					Amount       int `json:"amount"`
					Level        int `json:"level"`
					Progress     int `json:"progress"`
					NextLevelCap int `json:"nextLevelCap"`
				} `json:"ayena"`
			} `json:"world"`
		} `json:"experience"`
		Energy struct {
			Amount        int       `json:"amount"`
			WillRestoreAt time.Time `json:"willRestoreAt"`
			Cap           int       `json:"cap"`
		} `json:"energy"`
		Currency struct {
			Coin    int `json:"coin"`
			Mythium int `json:"mythium"`
		} `json:"currency"`
		Settings struct {
			Newsletters    bool `json:"newsletters"`
			SkipActivation bool `json:"skipActivation"`
		} `json:"settings"`
		ID         string `json:"_id"`
		Email      string `json:"email"`
		Username   string `json:"username"`
		EthAddress string `json:"ethAddress"`
		Location   struct {
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
			V         int       `json:"__v"`
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			Venari    []string  `json:"venari"`
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
		KnownVenari struct {
			CityOfTecta []string `json:"city-of-tecta"`
		} `json:"knownVenari"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		V         int       `json:"__v"`
		FactionID int       `json:"factionId"`
		Service   struct {
		} `json:"service"`
		Inventory struct {
		} `json:"inventory"`
		ActivatedAlphaPass struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Faction     struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
			Gilded               bool   `json:"gilded"`
			ObtainableBasePasses int    `json:"obtainableBasePasses"`
			Type                 string `json:"type"`
			Owner                struct {
				Experience struct {
					Account int `json:"account"`
					World   struct {
						Tecta int `json:"tecta"`
						Ayena int `json:"ayena"`
					} `json:"world"`
				} `json:"experience"`
				Energy struct {
					Amount        int       `json:"amount"`
					WillRestoreAt time.Time `json:"willRestoreAt"`
					Cap           int       `json:"cap"`
				} `json:"energy"`
				Currency struct {
					Coin    int `json:"coin"`
					Mythium int `json:"mythium"`
				} `json:"currency"`
				Settings struct {
					Newsletters    bool `json:"newsletters"`
					SkipActivation bool `json:"skipActivation"`
				} `json:"settings"`
				ID          string `json:"_id"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				EthAddress  string `json:"ethAddress"`
				Location    string `json:"location"`
				KnownVenari struct {
					CityOfTecta []string `json:"city-of-tecta"`
				} `json:"knownVenari"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
				V         int       `json:"__v"`
				FactionID int       `json:"factionId"`
				Service   struct {
				} `json:"service"`
				Inventory struct {
				} `json:"inventory"`
			} `json:"owner"`
			TokenID      int `json:"tokenId"`
			ActivatedFor struct {
				Experience struct {
					Account int `json:"account"`
					World   struct {
						Tecta int `json:"tecta"`
						Ayena int `json:"ayena"`
					} `json:"world"`
				} `json:"experience"`
				Energy struct {
					Amount        int       `json:"amount"`
					WillRestoreAt time.Time `json:"willRestoreAt"`
					Cap           int       `json:"cap"`
				} `json:"energy"`
				Currency struct {
					Coin    int `json:"coin"`
					Mythium int `json:"mythium"`
				} `json:"currency"`
				Settings struct {
					Newsletters    bool `json:"newsletters"`
					SkipActivation bool `json:"skipActivation"`
				} `json:"settings"`
				ID          string `json:"_id"`
				Email       string `json:"email"`
				Username    string `json:"username"`
				EthAddress  string `json:"ethAddress"`
				Location    string `json:"location"`
				KnownVenari struct {
					CityOfTecta []string `json:"city-of-tecta"`
				} `json:"knownVenari"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
				V         int       `json:"__v"`
				FactionID int       `json:"factionId"`
				Service   struct {
				} `json:"service"`
				Inventory struct {
				} `json:"inventory"`
			} `json:"activatedFor"`
			ActivatedAt      time.Time `json:"activatedAt"`
			AutoReactivation bool      `json:"autoReactivation"`
			LendPercentage   int       `json:"lendPercentage"`
			ShopDiscount     float64   `json:"shopDiscount"`
			Assets           struct {
				Image      string `json:"image"`
				Video      string `json:"video"`
				LargeVideo string `json:"largeVideo"`
			} `json:"assets"`
			EarningProgress struct {
				Amount      int       `json:"amount"`
				LastAddedAt time.Time `json:"lastAddedAt"`
			} `json:"earningProgress"`
		} `json:"activatedAlphaPass"`
		AlphaPasses []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Faction     struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"faction"`
			Gilded               bool      `json:"gilded"`
			ObtainableBasePasses int       `json:"obtainableBasePasses"`
			Type                 string    `json:"type"`
			TokenID              int       `json:"tokenId"`
			ActivatedAt          time.Time `json:"activatedAt"`
			AutoReactivation     bool      `json:"autoReactivation"`
			LendPercentage       int       `json:"lendPercentage"`
			ShopDiscount         float64   `json:"shopDiscount"`
			Assets               struct {
				Image      string `json:"image"`
				Video      string `json:"video"`
				LargeVideo string `json:"largeVideo"`
			} `json:"assets"`
			EarningProgress struct {
				Amount      int       `json:"amount"`
				LastAddedAt time.Time `json:"lastAddedAt"`
			} `json:"earningProgress"`
		} `json:"alphaPasses"`
	} `json:"data"`
}

func GetUser(client *requests.HttpClient) (*UserResponse, error) {

	req, _ := http.NewRequestWithContext(client.Context, "GET", "https://api.legendsofvenari.com/user", nil)
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

	var userRes UserResponse
	err = json.Unmarshal(resBytes, &userRes)
	if err != nil {
		return nil, err
	}

	return &userRes, nil
}
