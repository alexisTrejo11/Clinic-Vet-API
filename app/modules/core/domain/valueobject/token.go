package valueobject

type TokenType string

const (
	JWTAccessToken    TokenType = "jwt_access"
	JWTRefreshToken   TokenType = "jwt_refresh"
	TwoFAToken        TokenType = "2fa"
	OAuth2SecretToken TokenType = "oauth2_secret"
	ActivationToken   TokenType = "activation_token"
	VerificationToken TokenType = "verification_token"
)

type Token interface {
	Generate() (string, error)
	Validate(token string) (any, error)
	GetType() TokenType
	IsExpired() bool
}
