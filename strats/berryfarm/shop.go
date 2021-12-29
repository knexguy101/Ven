package berryfarm

import (
	"VenariBot/requests/user"
	"VenariBot/strats"
	"fmt"
	"strings"
	"time"
)

func Shop(vt *strats.VenariTask) string {

	fmt.Println("Buying Berries")

	//hardcode the berry id
	berryId := "61c4888ac120242ce8286d7a"

	items, err := user.BuyItem(berryId, 30, vt.Client)
	if err != nil {
		fmt.Println("Error buying berries")
		return "shop"
	}

	if strings.Contains(items.Message, "success") {
		return "expedition"
	}

	fmt.Printf("Failed to buy the berries: %s\n", items.Message)
	time.Sleep(10 * time.Second)
	return "shop"
}
