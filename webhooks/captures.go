package webhooks

import (
	"VenariBot/requests/expeditions/battle"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/webhook"
	"strconv"
)

var client *webhook.Client

func init() {
	client = webhook.NewClient("925800030974148649", "tzj4Q6xbRTQTAG2sMhowP1QA93m28k_MvTAvK_eT3bed1nWpIybcAtelqwNFsEGXGpDZ")
}

func SendCaptureWebhook(br *battle.StartBattleResponse, coinAmount int) {
	embed := discord.NewEmbedBuilder().
		SetTitle("Captured!").
		SetImage(br.Venari.Assets.Image).
		AddField("Name", br.Venari.Name, true).
		AddField("Tier", br.Venari.Tier, true).
		AddField("Reward", strconv.Itoa(coinAmount), false).
		SetColor(3145631).
		Build()

	client.CreateEmbeds([]discord.Embed{embed})
}

func SendFailedWebhook(br *battle.StartBattleResponse) {
	embed := discord.NewEmbedBuilder().
		SetTitle("Failed").
		SetImage(br.Venari.Assets.Image).
		AddField("Name", br.Venari.Name, true).
		AddField("Tier", br.Venari.Tier, true).
		SetColor(16730112).
		Build()

	client.CreateEmbeds([]discord.Embed{embed})
}
