package main

import (
	"github.com/skpr/cloudwatch-slack-lambda/internal/slack"
)

func main() {
	err := slack.PostMessage(slack.PostMessageParams{
		Webhooks: []string{
			"https://hooks.slack.com/services/xxxxxx/yyyyyyyyy/zzzzzzzzz",
		},
		Cluster:       "skpr-local",
		Project:       "test",
		Environment:   "dev",
		Description:   "This is a test description",
		Reason:        "This is a test reason",
		Dashboard:     "https://www.skpr.com.au",
		Documentation: "https://docs.skpr.io",
		Image:         "https://docs.skpr.io/assets/logo.svg",
	})
	if err != nil {
		panic(err)
	}
}
