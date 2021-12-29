package main

import (
	"VenariBot/requests"
	"VenariBot/strats"
	"VenariBot/strats/berryfarm"
)

func main() {
	berryfarm.StartTask(&strats.VenariTask{
		Client: requests.CreateNewClient(nil),
	})
}
