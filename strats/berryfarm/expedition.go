package berryfarm

import (
	"VenariBot/requests/expeditions"
	"VenariBot/requests/inventory"
	user2 "VenariBot/requests/user"
	"VenariBot/strats"
	"fmt"
	"time"
)

func Expedition(vt *strats.VenariTask) string {

	fmt.Println("Checking Energy")
	user, err := user2.GetUser(vt.Client)
	if err != nil {
		fmt.Println("Error getting user")
		return "expedition"
	}

	if user.Data.Energy.Amount < 2 {
		fmt.Println("Energy too low! sleeping...")
		time.Sleep(2 * time.Minute)
		return "expedition"
	}

	fmt.Println("Getting Expeditions")
	currentExp, err := expeditions.GetExpeditions("city-of-tecta", vt.Client)
	if err != nil {
		fmt.Println("Error getting expedition")
		return "expedition"
	}

	fmt.Println("Getting Inventory")
	currentInv, err := inventory.GetInventory(vt.Client)
	if err != nil {
		fmt.Println("Error getting inventory")
		return "expedition"
	}

	vt.BaitId = ""
	for _, v := range *currentInv {
		if v.Item.Name == "Oria" {
			vt.BaitId = v.ID
		} else if v.Item.Name == "Makeshift Rig" {
			vt.RigId = v.ID
		}
	}
	if vt.BaitId == "" {
		return "shop"
	}

	if len(*currentExp) > 0 {
		fmt.Println("Expedition found, starting")
		vt.ExpeditionId = (*currentExp)[0].ID
		vt.StartSearch = time.Now().Unix()
		return "search"
	}

	fmt.Println("Creating Expedition")
	fmt.Println(vt.BaitId)
	newExp, err := expeditions.CreateExpedition("city-of-tecta", vt.BaitId, vt.Client)
	if err != nil {
		fmt.Println("Error creating expedition")
		return "expedition"
	}

	vt.ExpeditionId = (*newExp)[0].ID
	vt.StartSearch = time.Now().Unix()
	return "search"
}
