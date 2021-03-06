package dolores_slack

import "github.com/nlopes/slack"

var (
	BotID                     string
	BotName                   string
	BotTextPrefixesForChannel []string
	DoloresAdminEmailIds      []string
	DbAdminEmailIds           []string
	API                       *slack.Client
)
