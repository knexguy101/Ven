package berryfarm

import "VenariBot/strats"

func StartTask(venariTask *strats.VenariTask) {
	status := "login"
	for {
		switch status {
		case "login":
			status = venariTask.Login("expedition")
			break
		case "expedition":
			status = Expedition(venariTask)
			break
		case "search":
			status = Search(venariTask)
			break
		case "battle":
			status = Battle(venariTask)
			break
		case "shop":
			status = Shop(venariTask)
			break
		}
	}
}
