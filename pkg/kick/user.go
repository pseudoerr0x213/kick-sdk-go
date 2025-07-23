package kick

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type GetUsersResponse struct {
	Data    []User `json:"data"`
	Message string `json:"message"`
}

func (c *Client) GetUsers(ctx context.Context, userIDs ...any) (*GetUsersResponse, error) {
	var resp GetUsersResponse

	q := url.Values{}
	q.Set("user_ids", joinIDs(userIDs...))
	path := "/public/v1/users?" + q.Encode()
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func joinIDs(ids ...any) string {
	res := make([]string, len(ids))
	for i, v := range ids {
		res[i] = fmt.Sprint(v)
	}
	return strings.Join(res, ",")
}
