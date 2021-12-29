package expeditions

import (
	"VenariBot/requests"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateExpeditionPayload struct {
	Bait string `json:"bait"`
	Area string `json:"area"`
}

func CreateExpedition(area, baitId string, client *requests.HttpClient) (*ExpeditionResponse, error) {

	resBytes, _ := json.Marshal(CreateExpeditionPayload{
		Area: area,
		Bait: baitId,
	})

	req, _ := http.NewRequestWithContext(client.Context, "POST", "https://api.legendsofvenari.com/expeditions", bytes.NewBuffer(resBytes))
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

	resBytes, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unknown response on csrf: %d", res.StatusCode)
	}

	var expRes ExpeditionResponse
	err = json.Unmarshal(resBytes, &expRes)
	if err != nil {
		return nil, err
	}

	return &expRes, nil
}