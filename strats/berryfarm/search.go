package berryfarm

import (
	"VenariBot/requests/expeditions"
	"VenariBot/strats"
	"fmt"
	"time"
)

func Search(vt *strats.VenariTask) string {

	catchAny := false
	if vt.StartSearch < time.Now().Add(-16 * time.Minute).Unix() {
		fmt.Println("Expedition Timeout, catching whatever")
		catchAny = true
	}

	fmt.Println("Searching for Venari")
	currentExp, err := expeditions.GetExpeditions("city-of-tecta", vt.Client)
	if err != nil {
		fmt.Println("Error getting expedition")
		return "search"
	}

	if len(*currentExp) <= 0 {
		fmt.Println("No expedition found, starting")
		return "expedition"
	}

	exp := (*currentExp)[0]

	for _, v := range exp.Spawns {
		if v.Venari.Name == "Vespille" && !catchAny {
			vt.VenariId = v.ID
			return "battle"
		} else if catchAny {
			vt.VenariId = v.ID
			return "battle"
		}
	}

	fmt.Println("No matching Venari found, sleeping...")
	time.Sleep(2 * time.Minute)

	return "search"
}
