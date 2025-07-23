package kick

import (
	"context"
)

type GetChannelsResponse struct {
	BannerPicture      string   `json:"banner_picture"`
	BroadcasterUserID  int64    `json:"broadcaster_user_id"`
	Category           Category `json:"category"`
	ChannelDescription string   `json:"channel_description"`
	Slug               string   `json:"slug"`
	Stream             Stream   `json:"stream"`
	StreamTitle        string   `json:"stream_title"`
}

type UpdateChannelRequest struct {
	CategoryID  int64  `json:"category_id"`
	StreamTitle string `json:"stream_title"`
}

func (c *Client) GetChannels(ctx context.Context) (*GetChannelsResponse, error) {
	var resp GetChannelsResponse
	path := "/public/v1/channels"
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateChannels(ctx context.Context, req UpdateChannelRequest) error {
	path := "/public/v1/channels"
	if err := c.doRequest(ctx, "PATCH", path, &req, nil); err != nil {
		return err
	}
	return nil
}
