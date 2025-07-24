package kick

import "context"

type GetPublicKeyResponse struct {
	Data    PublicKey `json:"data"`
	Message string    `json:"message"`
}

type PublicKey struct {
	PublicKey string `json:"public_key"`
}

func (c *Client) GetPublicKey(ctx context.Context) (*PublicKey, error) {
	var resp GetPublicKeyResponse
	path := "/public/v1/public-key"

	if err := c.doRequest(ctx, "GET", path, nil, &resp); err != nil {
		return nil, err
	}

	return &resp.Data, nil
}
