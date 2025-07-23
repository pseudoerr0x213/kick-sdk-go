package kick

import (
	"context"
	"fmt"
)

// const baseURL := "https://api.kick.com/public/v1/categories"

type GetCategoriesResponse struct {
	Data    []Category `json:"data"`
	Message string     `json:"message"`
}

type GetCategoryByIDResponse struct {
	Data    Category `json:"data"`
	Message string   `json:"message"`
}

func (c *Client) GetCategories(ctx context.Context) (*[]Category, error) {
	var resp GetCategoriesResponse
	if err := c.doRequest(ctx, "GET", "/public/v1/categories", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

func (c *Client) GetCategoryByID(ctx context.Context, categoryID int) (*Category, error) {
	var resp GetCategoryByIDResponse
	path := fmt.Sprintf("public/v1/categories/%d", categoryID)
	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err

	}
	return &resp.Data, nil
}
