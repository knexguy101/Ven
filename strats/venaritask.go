package strats

import (
	"Ven/requests"
	"Ven/requests/login"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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

func listenForEmailLink() string {
	app := fiber.New()
	defer app.Shutdown()

	tokenChan := make(chan string)

	app.Get("/sign-in", func(c *fiber.Ctx) error {
		token := utils.CopyString(c.Query("token"))
		tokenChan <- token
		return c.SendString("Loading into ven client...")
	})

	go func() {
		err := app.Listen(":3000")
		if err != nil {
			panic("Cannot listen on port 3000")
		}
	}()

	token, ok := <- tokenChan
	if !ok {
		panic("Did not receive a token")
	}
	return token
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
	token := listenForEmailLink()
	res, err := login.SubmitToken(token, vt.Client)
	if err != nil {
		return "login"
	}

	fmt.Println(fmt.Sprintf("Email: %s\nCurrency: %d", res.Data.Email, res.Data.Currency.Coin))

	return nextStatus
}
