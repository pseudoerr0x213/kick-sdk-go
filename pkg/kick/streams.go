package kick

type GetLiveStreamsResponse struct {
	Data    []LiveStream `json:"data"`
	Message string       `json:"message"`
}

type LiveStream struct {
	BroadcasterUserID int64    `json:"broadcaster_user_id"`
	Category          Category `json:"category"`
	ChannelID         int64    `json:"channel_id"`
	HasMatureContent  bool     `json:"has_mature_content"`
	Language          string   `json:"language"`
	Slug              string   `json:"slug"`
	StartedAt         string   `json:"started_at"`
	StreamTitle       string   `json:"stream_title"`
	Thumbnail         string   `json:"thumbnail"`
	ViewerCount       int64    `json:"viewer_count"`
}
