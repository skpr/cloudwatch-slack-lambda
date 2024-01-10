package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awscloudwatch "github.com/aws/aws-sdk-go-v2/service/cloudwatch"

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
func HandleLambdaEvent(ctx context.Context, event *cloudwatch.Event) error {
	log.Printf("Running Lambda (%s)\n", GitVersion)

	config, err := util.LoadConfig(".")
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	errs := config.Validate()
	if len(errs) > 0 {
		return fmt.Errorf("configuration error: %s", strings.Join(errs, "\n"))
	}

	slackClient, err := slack.NewClient(config.SlackWebhookURL)
	if err != nil {
		return fmt.Errorf("failed to create Slack client: %w", err)
	}

	cfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %v", err)
	}

	cloudwatchClient := awscloudwatch.NewFromConfig(cfg)

	err = run(ctx, cloudwatchClient, slackClient, event)
	if err != nil {
		return err
	}

	log.Println("Function complete")

	return nil
}

// Run will execute the core of the function.
func run(ctx context.Context, cloudwatchClient cloudwatch.ClientInterface, slackClient slack.ClientInterface, event *cloudwatch.Event) error {
	if event.AlarmARN == "" {
		return fmt.Errorf("alarm ARN is required")
	}

	alarm, err := cloudwatchClient.ListTagsForResource(ctx, &awscloudwatch.ListTagsForResourceInput{
		ResourceARN: aws.String(event.AlarmARN),
	})
	if err != nil {
		return fmt.Errorf("failed to list tags for resource: %w", err)
	}

	err = slackClient.PostMessage(slack.PostMessageParams{
		Cluster:       cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyCluster),
		Project:       cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyProject),
		Environment:   cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyEnvironment),
		Description:   event.AlarmData.Configuration.Description,
		Reason:        event.AlarmData.State.Reason,
		Dashboard:     cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyLinkDashboard),
		Documentation: cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyLinkDocumentation),
		Image:         cloudwatch.GetValueFromTag(alarm.Tags, cloudwatch.TagKeyAssetIcon),
	})
	if err != nil {
		return fmt.Errorf("failed to post Slack message: %w", err)
	}

	return nil
}
