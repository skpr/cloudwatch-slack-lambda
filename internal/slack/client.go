package slack

// ClientInterface is used to interact with Slack.
type ClientInterface interface {
	PostMessage(params PostMessageParams) error
}

// Client is used to interact with Slack.
type Client struct {
	webhooks []string
}

// NewClient for interacting with Slack.
func NewClient(webhooks []string) (*Client, error) {
	client := &Client{
		webhooks: webhooks,
	}

	return client, nil
}
