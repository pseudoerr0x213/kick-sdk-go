package kick

type Ban struct{}

type Category struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

type User struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profile_picture"`
	UserID         int64  `json:"user_id"`
}
type Stream struct {
	IsLive      bool   `json:"is_live"`
	IsMature    bool   `json:"is_mature"`
	Key         string `json:"key"`
	Language    string `json:"language"`
	StartTime   string `json:"start_time"`
	Thumbnail   string `json:"thumbnail"`
	URL         string `json:"url"`
	ViewerCount int64  `json:"viewer_count"`
}
