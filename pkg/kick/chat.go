package kick

import (
	"context"
	"fmt"
)

// type field can only have two values: "user" or "bot". See: https://docs.kick.com/apis/chat
type ChatMessageType string

const (
	ChatMessageUser ChatMessageType = "user"
	ChatMessageBot  ChatMessageType = "bot"
)

type SendChatRequest struct {
	BroadcasterUserID int64           `json:"broadcaster_user_id"`
	Content           string          `json:"content"`
	ReplyToMessageID  string          `json:"reply_to_message_id"`
	Type              ChatMessageType `json:"type"`
}
type SendChatResponse struct {
	IsSent    bool   `json:"is_sent"`
	MessageID string `json:"message_id"`
}

func (c *Client) PostChatMessage(
	ctx context.Context,
	broadcasterUserID int64,
	content string,
	reply_to_msg_id string, msgType string) (*SendChatResponse, error) {
	var req SendChatRequest

	switch msgType {
	case "user":
		msgType = ChatMessageType(ChatMessageUser)

		if msgType == ChatMessageType(ChatMessageUser) && broadcasterUserID == 0 {
			return nil, fmt.Errorf("broadcaster userID is required")
		}

	case "bot":
		msgType = ChatMessageBot(ChatMessageBot)

	default:
		return nil, fmt.Errorf("the msgType should be user or bot")
	}

	return req, nil
}

type SendChatRequestOption func(*SendChatRequest)

func NewSendChatRequest(broadcasterUserID int64, content string, reply_to_msg_id string, msgType string) (*SendChatRequest, error) {
	switch msgType {
	case "user":
		msgType = ChatMessageType(ChatMessageUser)
	case "bot":
		msgType = ChatMessageBot(ChatMessageBot)

	default:
		return nil, fmt.Errorf("the msgType can be user or bot")
	}

	return &SendChatRequest{
		BroadcasterUserID: broadcasterUserID,
		Content:           content,
		ReplyToMessageID:  reply_to_msg_id,
	}

}

// constructor for sending chat messages from user
func UserSendChatRequest(broadcasterUserID int64, content string, replyToMsg string) *SendChatRequestOption {

}
