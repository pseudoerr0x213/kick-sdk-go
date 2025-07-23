package kick

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type Duration int64

var (
	ErrTimeout         error = errors.New("timeout value should be more than 1 and less than 10080")
	ErrNoRequiredField error = errors.New("no required field provided")
)

const (
	MinDuration Duration = 1
	MaxDuration Duration = 10080
)

const (
	modPath string = "/public/v1/moderation/bans"
)

type CreateBanRequest struct {
	BroadcasterUserID int64    `json:"broadcaster_user_id"`
	Duration          Duration `json:"duration,omitempty"` // optional
	Reason            string   `json:"reason,omitempty"`   // optional
	UserID            int64    `json:"user_id"`
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

func (c *Client) PostBan(ctx context.Context, broadcasterID int64, userID int64, opts ...CreateBanRequestOption) (*CreateBanResponse, error) {
	req, err := NewCreateBanRequest(broadcasterID, userID, opts...)
	if err != nil {
		return nil, err
	}

	var resp CreateBanResponse
	if err := c.doRequest(ctx, "POST", modPath, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) DeleteBan(ctx context.Context, broadcasterID int64, userID int64) (*DeleteBanResponse, error) {
	req, err := NewDeleteBanRequest(broadcasterID, userID)
	if err != nil {
		return nil, err
	}

	var resp DeleteBanResponse
	if err := c.doRequest(ctx, "DELETE", modPath, req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func NewCreateBanRequest(broadcasterUserID int64, userID int64, opts ...CreateBanRequestOption) (*CreateBanRequest, error) {
	if userID == 0 || broadcasterUserID == 0 {
		return nil, ErrNoRequiredField
	}

	req := &CreateBanRequest{
		BroadcasterUserID: broadcasterUserID,
		UserID:            userID,
	}

	for _, opt := range opts {
		opt(req)
	}

	return req, nil
}

func NewDeleteBanRequest(broadcasterUserID int64, userID int64) (*DeleteBanRequest, error) {
	if userID == 0 || broadcasterUserID == 0 {
		return nil, ErrNoRequiredField
	}

	return &DeleteBanRequest{
		BroadcasterUserID: broadcasterUserID,
		UserID:            userID,
	}, nil
}

// options
type CreateBanRequestOption func(*CreateBanRequest)

// using this option will timeout the user
// if you want to ban him instead, don't use this
// see: https://docs.kick.com/apis/moderation#post-moderation-bans
func Timeout(duration int64) CreateBanRequestOption {
	return func(cbr *CreateBanRequest) {
		cbr.Duration, ErrTimeout = ValidateDuration(duration)
		if ErrTimeout != nil {
			log.Println(ErrTimeout)
			return
		}
	}
}

func WithReason(reason string) CreateBanRequestOption {
	return func(cbr *CreateBanRequest) {
		cbr.Reason = reason
	}
}

// helper functions
func ValidateDuration(val int64) (Duration, error) {
	if val < int64(MinDuration) || val > int64(MaxDuration) {
		return 0, fmt.Errorf("duration must be between %d and %d", MinDuration, MaxDuration)
	}
	return Duration(val), nil
}
