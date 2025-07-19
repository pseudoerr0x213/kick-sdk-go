package kicksdk

type SendChatRequest struct {
	BroadcasterUserID int64  `json:"broadcaster_user_id"`
	Content           string `json:"content"`
	ReplyToMessageID  string `json:"reply_to_message_id"`
	Type              string `json:"type"`
}

type SendChatResponse struct {
	IsSent    bool   `json:"is_sent"`
	MessageID string `json:"message_id"`
}
