package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// PostMessageParams are the parameters required to post a message to Slack.
type PostMessageParams struct {
	// Metadata.
	Cluster     string
	Project     string
	Environment string

	// Details.
	Description string
	Reason      string

	// Actions.
	Dashboard     string
	Documentation string

	// Image which will be applied to this message.
	Image string
}

// Validate the parameters.
func (p PostMessageParams) Validate() error {
	var errs []error

	if p.Cluster == "" {
		errs = append(errs, fmt.Errorf("cluster is required"))
	}

	if p.Project == "" {
		errs = append(errs, fmt.Errorf("project is required"))
	}

	if p.Environment == "" {
		errs = append(errs, fmt.Errorf("environment is required"))
	}

	if p.Description == "" {
		errs = append(errs, fmt.Errorf("description is required"))
	}

	if p.Reason == "" {
		errs = append(errs, fmt.Errorf("reason is required"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// PostMessage to Slack channel.
func (c *Client) PostMessage(params PostMessageParams) error {
	if err := params.Validate(); err != nil {
		return fmt.Errorf("invalid parameters: %w", err)
	}

	var message Message

	// Context which allows the developer to understand what project/environment is affected.
	// This is intentionally ordered as Environment/Project/Cluster because as a developer, I would expect someone
	// to review the message in that order.
	//   * Is it production?
	//   * Which project?
	//   * Which cluster?
	message.Blocks = append(message.Blocks, BlockContext{
		Type: BlockTypeContext,
		Elements: []BlockContextElement{
			{
				Type: BlockElementTypeMarkdown,
				Text: fmt.Sprintf("*Environment* = %s", params.Environment),
			},
			{
				Type: BlockElementTypeMarkdown,
				Text: fmt.Sprintf("*Project* = %s", params.Project),
			},
			{
				Type: BlockElementTypeMarkdown,
				Text: fmt.Sprintf("*Cluster* = %s", params.Cluster),
			},
		},
	})

	// Separate the context from the content.
	message.Blocks = append(message.Blocks, BlockDivider{
		Type: BlockTypeDivider,
	})

	// Details of the alarm.
	details := BlockSection{
		Type: BlockTypeSection,
		Text: BlockSectionText{
			Type: BlockTextTypeMarkdown,
			Text: fmt.Sprintf("*%s*\n\n_%s_", params.Description, params.Reason),
		},
	}

	if params.Image != "" {
		details.Accessory = &BlockSectionAccessory{
			Type:     BlockElementTypeImage,
			ImageURL: params.Image,
			AltText:  "Identifier for the Slack message",
		}
	}

	message.Blocks = append(message.Blocks, details)

	// Links which can be used to action message eg. Go to this dashboard or this documentation.
	var links []string

	if params.Dashboard != "" {
		links = append(links, fmt.Sprintf("<%s|:skpr_dashboard: Go to Dashboard>", params.Dashboard))
	}

	if params.Documentation != "" {
		links = append(links, fmt.Sprintf("<%s|:skpr_documentation: Go to Documentation>", params.Documentation))
	}

	if len(links) > 0 {
		message.Blocks = append(message.Blocks, BlockSection{
			Type: BlockTypeSection,
			Text: BlockSectionText{
				Type: BlockTextTypeMarkdown,
				Text: strings.Join(links, "\t"),
			},
		})
	}

	request, err := json.Marshal(message)
	if err != nil {
		return err
	}

	fmt.Println(string(request))

	for _, webhook := range c.webhooks {
		req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(request))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)

		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("returned status code: %d", resp.StatusCode)
		}
	}

	return nil
}
