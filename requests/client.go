package requests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	Tr *http.Transport
	Cookies map[string]*http.Cookie
	Context context.Context
	Cancel func()
}

func CreateNewClient(proxy *url.URL) *HttpClient {
	client := HttpClient {
		Tr: &http.Transport{
			ForceAttemptHTTP2: true,
			DisableKeepAlives: true,
			Proxy: http.ProxyURL(proxy),
		},
		Cookies: make(map[string]*http.Cookie),
	}
	client.Context, client.Cancel = context.WithCancel(context.Background())
	return &client
}

func (hc *HttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Cookie", hc.GetCookieHeader())

	res, err := hc.Tr.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	//dont close btw

	for _, v := range res.Cookies() {
		hc.Cookies[v.Name] = v
	}

	return res, nil
}

func (hc *HttpClient) GetCookieHeader() string {
	var cookies []string
	for _, v := range hc.Cookies {
		cookies = append(cookies, fmt.Sprintf("%s=%s", v.Name, v.Value))
	}
	return strings.Join(cookies, "; ")
}

func (hc *HttpClient) GetXSRF() string {
	val, ok := hc.Cookies["XSRF-TOKEN"]
	if !ok {
		return ""
	}
	return val.Value
}
