package main

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	cloudwatchtypes "github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/stretchr/testify/assert"

	"github.com/skpr/cloudwatch-slack-lambda/internal/cloudwatch"
	"github.com/skpr/cloudwatch-slack-lambda/internal/slack"
)

func TestRun(t *testing.T) {
	event := &cloudwatch.Event{
		AlarmARN: "arn:aws:cloudwatch:ap-southeast-2:xxxxxxxxx:alarm:xxxxxxxxx",
		AlarmData: cloudwatch.AlarmData{
			AlarmName: "test-alarm",
			State: cloudwatch.AlarmDataState{
				Reason: "This is a test reason",
			},
			Configuration: cloudwatch.AlarmDataConfiguration{
				Description: "This is a test description",
			},
		},
	}

	cloudwatchMock := &cloudwatch.MockClient{
		Tags: []cloudwatchtypes.Tag{
			{
				Key:   aws.String(cloudwatch.TagKeyCluster),
				Value: aws.String("skpr-test"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyProject),
				Value: aws.String("test"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyEnvironment),
				Value: aws.String("dev"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyInstance),
				Value: aws.String("nonprod"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyLinkDashboard),
				Value: aws.String("https://www.skpr.io"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyLinkDocumentation),
				Value: aws.String("https://docs.skpr.io"),
			},
			{
				Key:   aws.String(cloudwatch.TagKeyAssetIcon),
				Value: aws.String("https://docs.skpr.io/icon.png"),
			},
		},
	}

	slackMock := &slack.MockClient{}

	err := run(context.TODO(), cloudwatchMock, slackMock, event)
	assert.NoError(t, err)

	want := slack.PostMessageParams{
		Cluster:       "skpr-test",
		Project:       "test",
		Environment:   "dev",
		Instance:      "nonprod",
		Description:   "This is a test description",
		Reason:        "This is a test reason",
		Dashboard:     "https://www.skpr.io",
		Documentation: "https://docs.skpr.io",
		Image:         "https://docs.skpr.io/icon.png",
	}

	assert.Equal(t, want, slackMock.PostMessageParams)
}
