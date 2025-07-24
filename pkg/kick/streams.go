package kick

import "context"

const path = "/public/v1/livestreams"

type Sort string

const (
	ByViewerCount Sort = "viewer_count"
	ByStartedAt   Sort = "started_at"
)

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

func (c *Client) GetLiveStreams(ctx context.Context, broadcasterUserID int64) (*[]LiveStream, error) {
	var resp GetLiveStreamsResponse

	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
