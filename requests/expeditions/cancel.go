package expeditions

import (
	"Ven/requests"
	"fmt"
	"net/http"
)

func CancelExpedition(expeditionId string, client *requests.HttpClient) error {

	req, _ := http.NewRequestWithContext(client.Context, "DELETE", fmt.Sprintf("https://api.legendsofvenari.com/expeditions/%s", expeditionId), nil)
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
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("unknown response on csrf: %d", res.StatusCode)
	}

	return nil
}
