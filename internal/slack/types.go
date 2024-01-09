package slack

type BlockType string

const (
	BlockTypeContext BlockType = "context"
	BlockTypeSection BlockType = "section"
	BlockTypeDivider BlockType = "divider"
	BlockTypeActions BlockType = "actions"
)

type BlockElementType string

const (
	// BlockElementTypeMarkdown used to declare a block as markdown.
	BlockElementTypeMarkdown BlockElementType = "mrkdwn"
	BlockElementTypeButton   BlockElementType = "button"
	BlockElementTypeImage    BlockElementType = "image"
)

type BlockTextType string

const (
	// BlockTextTypeMarkdown used to declare a block as markdown.
	BlockTextTypeMarkdown BlockTextType = "mrkdwn"
	// BlockTextTypePlainText used to declare a block as plain text.
	BlockTextTypePlainText BlockTextType = "plain_text"
)

type Message struct {
	Blocks []interface{} `json:"blocks"`
}

type BlockContext struct {
	Type     BlockType             `json:"type"`
	Elements []BlockContextElement `json:"elements"`
}

type BlockContextElement struct {
	Type BlockElementType `json:"type"`
	Text string           `json:"text"`
}

type BlockDivider struct {
	Type BlockType `json:"type"`
}

type BlockSection struct {
	Type      BlockType              `json:"type"`
	Text      BlockSectionText       `json:"text"`
	Accessory *BlockSectionAccessory `json:"accessory,omitempty"`
}

type BlockSectionAccessory struct {
	Type     BlockElementType `json:"type"`
	ImageURL string           `json:"image_url"`
	AltText  string           `json:"alt_text"`
}

type BlockSectionText struct {
	Type BlockTextType `json:"type"`
	Text string        `json:"text"`
}
