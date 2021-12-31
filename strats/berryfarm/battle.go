package berryfarm

import (
	"VenariBot/requests/expeditions/battle"
	"VenariBot/strats"
	"VenariBot/webhooks"
	"fmt"
)

func Battle(vt *strats.VenariTask) string {

	fmt.Println("Starting Battle")
	start, err := battle.StartBattle(vt.ExpeditionId, vt.VenariId, vt.BaitId, vt.RigId, vt.Client)
	if err != nil {
		fmt.Println("Error starting battle")
		return "battle"
	}

	fmt.Println(fmt.Sprintf("Started Battle: %s", start.Venari.Name))

	for i := 0; i < 3; i++ {

		fmt.Printf("Fighting [%d]\n", i)
		play, err := battle.BattleAction(vt.ExpeditionId, vt.BaitId, vt.RigId, battle.Play, vt.Client)
		if err != nil {
			fmt.Println("Error playing")
			return "battle"
		}

		fmt.Println(fmt.Sprintf("[%s] %v %s", play.Venari, play.Success, play.CaptureChance))
	}

	capture, err := battle.BattleCatch(vt.ExpeditionId, vt.Client)
	if err != nil {
		fmt.Println("Error attempting capture")
		return "battle"
	}

	if len(capture.Rewards) > 0 {
		webhooks.SendCaptureWebhook(start, capture.Rewards[0].Amount)
		fmt.Println(fmt.Sprintf("Captured! %d %s", capture.Rewards[0].Amount, capture.Rewards[0].Type))
	} else {
		webhooks.SendFailedWebhook(start)
		fmt.Println("Failed to capture")
	}

	return "expedition"
}
