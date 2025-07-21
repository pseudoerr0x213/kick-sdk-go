package kick

type CreateBanRequest struct {
	BroadcasterUserID int64  `json:"broadcaster_user_id"`
	Duration          int64  `json:"duration"`
	Reason            string `json:"reason"`
	UserID            int64  `json:"user_id"`
}

type CreateBanResponse struct {
	Data    Ban    `json:"data"`
	Message string `json:"message"`
}

type DeleteBanRequest struct {
	BroadcasterUserID int64 `json:"broadcaster_user_id"`
	UserID            int64 `json:"user_id"`
}

type DeleteBanResponse struct {
	Data    Ban    `json:"data"`
	Message string `json:"message"`
}
