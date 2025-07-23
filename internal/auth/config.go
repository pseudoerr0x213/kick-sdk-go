package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// AuthConfig holds config for both App and User OAuth flows.
type AuthConfig struct {
	clientID     string
	clientSecret string

	// optional for user flow
	redirectURI  string
	scopes       []string
	codeVerifier string
	state        string

	// internal flag: "app" or "user"
	flow string

	// tokens + expiry
	accessToken  string
	refreshToken string
	expiresAt    time.Time
}

type AuthOption func(*AuthConfig)

// WithUserFlow enables PKCE user-auth flow
func WithUserFlow(
	redirectURI, codeVerifier, state string,
	scopes []string,
) AuthOption {
	return func(cfg *AuthConfig) {
		cfg.flow = "user"
		cfg.redirectURI = redirectURI
		cfg.codeVerifier = codeVerifier
		cfg.state = state
		cfg.scopes = scopes
	}
}

// NewAuthConfig creates a default AuthConfig for app-only flow or user flow if options provided.
func NewAuthConfig(id, secret string, opts ...AuthOption) *AuthConfig {
	cfg := &AuthConfig{
		clientID:     id,
		clientSecret: secret,
		flow:         "app",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// AuthProvider defines how to retrieve a bearer token.
type AuthProvider interface {
	Token(ctx context.Context) (string, error)
}

// Ensure AuthConfig implements AuthProvider
var _ AuthProvider = (*AuthConfig)(nil)

// Token returns a valid token depending on the configured flow.
func (cfg *AuthConfig) Token(ctx context.Context) (string, error) {
	switch cfg.flow {
	case "app":
		return cfg.appToken(ctx)
	case "user":
		return cfg.userToken(ctx)
	default:
		return "", fmt.Errorf("unknown auth flow: %s", cfg.flow)
	}
}

// appToken implements client credentials flow.
func (cfg *AuthConfig) appToken(ctx context.Context) (string, error) {
	if time.Now().Before(cfg.expiresAt) && cfg.accessToken != "" {
		return cfg.accessToken, nil
	}
	// build form
	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {cfg.clientID},
		"client_secret": {cfg.clientSecret},
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "https://id.kick.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("app token endpoint: %s", resp.Status)
	}

	var out struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}

	cfg.accessToken = out.AccessToken
	cfg.expiresAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	return cfg.accessToken, nil
}

// AuthURL returns the URL to redirect users for authorization (user flow).
func (cfg *AuthConfig) AuthURL() (string, error) {
	if cfg.flow != "user" {
		return "", fmt.Errorf("AuthURL only valid for user flow")
	}
	base := "https://id.kick.com/oauth/authorize"
	params := url.Values{
		"response_type":         {"code"},
		"client_id":             {cfg.clientID},
		"redirect_uri":          {cfg.redirectURI},
		"scope":                 {strings.Join(cfg.scopes, " ")},
		"state":                 {cfg.state},
		"code_challenge":        {cfg.codeVerifier},
		"code_challenge_method": {"S256"},
	}
	return base + "?" + params.Encode(), nil
}

// ExchangeCode exchanges the authorization code for tokens (user flow).
func (cfg *AuthConfig) ExchangeCode(ctx context.Context, code string) error {
	if cfg.flow != "user" {
		return fmt.Errorf("ExchangeCode only valid for user flow")
	}
	form := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {cfg.clientID},
		"client_secret": {cfg.clientSecret},
		"redirect_uri":  {cfg.redirectURI},
		"code_verifier": {cfg.codeVerifier},
		"code":          {code},
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "https://id.kick.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("exchange code: %s", resp.Status)
	}

	var out struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return err
	}

	cfg.accessToken = out.AccessToken
	cfg.refreshToken = out.RefreshToken
	cfg.expiresAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	return nil
}

// userToken implements refresh logic for user flow.
func (cfg *AuthConfig) userToken(ctx context.Context) (string, error) {
	if time.Now().Before(cfg.expiresAt) && cfg.accessToken != "" {
		return cfg.accessToken, nil
	}
	// refresh
	form := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {cfg.clientID},
		"client_secret": {cfg.clientSecret},
		"refresh_token": {cfg.refreshToken},
	}
	req, err := http.NewRequestWithContext(ctx, "POST", "https://id.kick.com/oauth/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("refresh token: %s", resp.Status)
	}

	var out struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}

	cfg.accessToken = out.AccessToken
	cfg.refreshToken = out.RefreshToken
	cfg.expiresAt = time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	return cfg.accessToken, nil
}
