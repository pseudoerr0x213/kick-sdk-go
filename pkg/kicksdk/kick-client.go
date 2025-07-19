package kicksdk

type KickClient struct {
	AppAuthConfig
	UserAuthConfig
}

type AppAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type UserAuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string

	// optinal helpers
	CodeVerifier string
	State        string
}

var (
	AuthURL   = "https://id.kick.com/oauth/authorize"
	TokenURL  = "https://id.kick.com/oauth/token"
	RevokeURL = "https://id.kick.com/oauth/revoke"
)

func NewKickClient()
