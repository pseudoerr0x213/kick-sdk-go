package kick

import (
	"context"
	"time"

	"github.com/pseudoerr/kick-sdk-go/internal/auth"
	"github.com/pseudoerr/kick-sdk-go/internal/http"
)

type Client struct {
	http *http.Client
	auth auth.AuthConfig
}

// Option lets you tweak timeouts, base URL, etc.
type Option func(*Client)

func WithHTTPTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.http = http.New(c.http.BaseURL, d)
	}
}

func NewClient(clientID, clientSecret string, opts ...Option) *Client {
	httpc := http.New("https://api.kick.com", 10*time.Second)
	auth := auth.NewAuthConfig(clientID, clientSecret)

	c := &Client{http: httpc, auth: &auth}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) doRequest(ctx context.Context, method, path string, body, out any) error {
	req, err := c.http.NewRequest(ctx, method, path, body)
	if err != nil {
		return err
	}

	token, err := c.auth.Token(ctx)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return c.http.Do(req, out)
}
