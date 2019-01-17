package util

const (
	MessageType            = iota
	EditedMessageType      = iota
	ChannelPostType        = iota
	EditedChannelPostType  = iota
	InlineQueryType        = iota
	ChosenInlineResultType = iota
	CallbackQueryType      = iota
	ShippingQueryType      = iota
	PreCheckoutQueryType   = iota
	OtherType              = -1
)

// Chat is a chat object of Telegram.
type Chat struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

// TgMessage is a message object of Telegram.
type TgMessage struct {
	MessageID int    `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

// TgUpdate is an update object of Telegram.
type TgUpdate struct {
	UpdateID      int        `json:"update_id"`
	Message       *TgMessage `json:"message,omitempty"`
	Editedmessage *TgMessage `json:"edited_message,omitempty"`
	InlineQuery   string     `json:"inline_query,omitempty"`
}

// UpdateType returns type of update JSON.
func (p TgUpdate) UpdateType() int {
	switch {
	case p.Message != nil:
		return MessageType
	case p.Editedmessage != nil:
		return EditedMessageType
	default:
		return OtherType
	}
}
