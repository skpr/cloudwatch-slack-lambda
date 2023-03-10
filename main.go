package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/skpr/cloudwatch-slack-lambda/internal/cloudwatch"
	"github.com/skpr/cloudwatch-slack-lambda/internal/slack"
	"github.com/skpr/cloudwatch-slack-lambda/internal/util"
)

var (
	// GitVersion overridden at build time by:
	//   -ldflags="-X main.GitVersion=${VERSION}"
	GitVersion string
)

func main() {
	lambda.Start(HandleLambdaEvent)
}

// HandleLambdaEvent will respond to a CloudWatch Alarm, check for rate limited IP addresses and send a message to Slack.
func HandleLambdaEvent(ctx context.Context, e events.CloudWatchEvent) error {
	log.Printf("Running Lambda (%s)\n", GitVersion)

	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	errs := config.Validate()
	if len(errs) > 0 {
		return fmt.Errorf("configuration error: %s", strings.Join(errs, "\n"))
	}

	log.Println("Inspecting event")

	var detail cloudwatch.EventDetail

	if err := json.Unmarshal(e.Detail, &detail); err != nil {
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	if !cloudwatch.HasBecomeAlarm(detail) {
		log.Printf("Skipping. Event did not become alarm (previous state = %s, current state = %s)\n",
			detail.PreviousState.Value,
			detail.State.Value)
		return nil
	}

	log.Println("Sending Slack message")

	err = slack.PostMessage(config, detail.AlarmName, detail.State.Reason)
	if err != nil {
		return fmt.Errorf("failed to post Slack message: %w", err)
	}

	log.Println("Function complete")

	return nil
}
