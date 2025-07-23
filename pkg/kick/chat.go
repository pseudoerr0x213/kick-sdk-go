package kick

import (
	"context"
	"errors"
)

var (
	ErrMissingContent       = errors.New("content is required")
	ErrMissingBroadcasterID = errors.New("broadcaster_user_id is required")
	ErrInvalidMessageType   = errors.New("type must be \"user\" or \"bot\"")
)

// type field can only have two values: "user" or "bot". See: https://docs.kick.com/apis/chat
type ChatMessageType string

const (
	ChatMessageUser ChatMessageType = "user"
	ChatMessageBot  ChatMessageType = "bot"
)

type SendChatRequest struct {
	BroadcasterUserID int64           `json:"broadcaster_user_id,omitempty"` //optional
	Content           string          `json:"content"`
	ReplyToMessageID  string          `json:"reply_to_message_id,omitempty"` //optional
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
	msgType ChatMessageType,
	opts ...SendChatRequestOption,
) (*SendChatResponse, error) {

	req, err := NewSendChatRequest(content, msgType, opts...)
	if err != nil {
		return nil, err
	}

	req.BroadcasterUserID = broadcasterUserID

	if req.Content == "" {
		return nil, ErrMissingContent
	}

	path := "/public/v1/chat/message"
	var resp SendChatResponse
	if err := c.doRequest(ctx, "POST", path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SendChatRequestOption func(*SendChatRequest)

func NewSendChatRequest(
	content string,
	msgType ChatMessageType,
	opts ...SendChatRequestOption,
) (*SendChatRequest, error) {
	if content == "" {
		return nil, ErrMissingContent
	}
	if msgType != ChatMessageUser && msgType != ChatMessageBot {
		return nil, ErrInvalidMessageType
	}
	req := &SendChatRequest{
		Content: content,
		Type:    msgType,
	}

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

// use this if you want to send a reply to
func WithReplyToMessageID(id string) SendChatRequestOption {
	return func(r *SendChatRequest) {
		r.ReplyToMessageID = id
	}
}
