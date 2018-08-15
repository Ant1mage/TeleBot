package teleBot

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// NewMessage creates a new Message.
//
// chatID is where to send it, text is the message text.
func NewMessage(chatID int64, messageType int) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		DisableWebPagePreview: false,
		MessageType:           messageType,
	}
}

func Send(msg MessageConfig) (Message, error) {

	params := url.Values{}
	// TODO: 更新Type
	switch msg.MessageType {
	case 0:
		params.Add("chat_id", fmt.Sprintf("%d", msg.ChatID))
		params.Add("text", msg.Text)
	case 1:
		params.Add("chat_id", fmt.Sprintf("%d", msg.ChatID))
		params.Add("photo", msg.Photo)
	default:
		break
	}

	endpoint := Endpoint[msg.MessageType]

	message, err := makeMessageRequest(endpoint, params)

	if err != nil {
		return Message{}, err
	}

	return message, nil
}

func makeMessageRequest(endpoint string, params url.Values) (Message, error) {
	resp, err := MakeRequest(endpoint, params)
	if err != nil {
		return Message{}, err
	}

	var message Message
	json.Unmarshal(resp.Result, &message)
	return message, nil
}
