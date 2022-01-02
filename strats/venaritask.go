package strats

import (
	"Ven/requests"
	"Ven/requests/login"
	"fmt"
)

type VenariTask struct {
	Client *requests.HttpClient

	ExpeditionId string
	BaitId string
	RigId string
	VenariId string

	StartSearch int64
}

func getConsoleInput() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return ""
	}
	return input
}

func (vt *VenariTask) Login(nextStatus string) string {
	err := login.GetCSRF(vt.Client)
	if err != nil {
		return "login"
	}

	fmt.Println("Enter your email")
	email := getConsoleInput()
	err = login.SubmitEmail(email, vt.Client)
	if err != nil {
		return "login"
	}

	fmt.Println("Enter your auth token")
	token := getConsoleInput()
	res, err := login.SubmitToken(token, vt.Client)
	if err != nil {
		return "login"
	}

	fmt.Println(fmt.Sprintf("Email: %s\nCurrency: %d", res.Data.Email, res.Data.Currency.Coin))

	return nextStatus
}
