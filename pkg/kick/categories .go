package kick

import (
	"context"
)

// const baseURL := "https://api.kick.com/public/v1/categories"

type GetCategoriesResponse struct {
	Data    []Category `json:"data"`
	Message string     `json:"message"`
}

func (c *Client) GetCategories(ctx context.Context) ([]Category, error) {
	var resp GetCategoriesResponse
	if err := c.doRequest(ctx, "GET", "/public/v1/categories", nil, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
