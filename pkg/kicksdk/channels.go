package kicksdk

type ChannelResponse struct {
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
