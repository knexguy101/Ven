package battle

import (
	"VenariBot/requests"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BattleActionType string
const (
	Play BattleActionType = "play"
	Fight BattleActionType = "fight"
	Feed BattleActionType = "feed"
)

type BattleActionPayload struct {
	Action BattleActionType `json:"action"`
	Rig    string `json:"rig"`
	Bait   string `json:"bait"`
}

type BattleActionResponse struct {
	CaptureChance string `json:"captureChance"`
	ActionNumber  int    `json:"actionNumber"`
	Venari        string `json:"venari"`
	Message       string `json:"message"`
	Success       bool   `json:"success"`
}

func BattleAction(expeditionId, baitId, rigId string, action BattleActionType, client *requests.HttpClient) (*BattleActionResponse, error) {

	resBytes, _ := json.Marshal(BattleActionPayload{
		Action: action,
		Rig: rigId,
		Bait: baitId,
	})

	req, _ := http.NewRequestWithContext(client.Context, "POST", fmt.Sprintf("https://api.legendsofvenari.com/expeditions/%s/action", expeditionId), bytes.NewBuffer(resBytes))
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

	resBytes, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var expRes BattleActionResponse
	err = json.Unmarshal(resBytes, &expRes)
	if err != nil {
		return nil, err
	}

	return &expRes, nil
}
