package dolores_slack

import (
	"strings"

	"github.com/abhishekkr/gol/golconfig"
	"github.com/nlopes/slack"
)

func LoopRTMEvents(config golconfig.FlatConfig) {
	overrideMessagesFromEnv()

	BotID = config["slack-bot-name"]
	DoloresAdminEmailIds = strings.Fields(config["admin-emails"])
	DbAdminEmailIds = strings.Fields(config["db-admin-emails"])
	API = AuthenticatedApi(config["slack-bot-api-token"])
	if config["slack-debug-mode"] == "true" {
		API.SetDebug(true)
	} else {
		API.SetDebug(false)
	}

	rtm := API.NewRTM()
	go rtm.ManageConnection() // spawn slack bot

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.HelloEvent:
				HelloEvent(ev)

			case *slack.ConnectedEvent:
				ConnectedEvent(ev)

			case *slack.MessageEvent:
				MessageEvent(ev)

			case *slack.PresenceChangeEvent:
				PresenceChangeEvent(ev)

			case *slack.LatencyReport:
				LatencyReport(ev)

			case *slack.RTMError:
				RTMError(ev)

			case *slack.InvalidAuthEvent:
				InvalidAuthEvent(ev)
				panic("invalid slack api token")

			default:
				DefaultEvent(msg)
			}
		}
	}
}
