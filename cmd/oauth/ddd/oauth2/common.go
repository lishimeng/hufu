package oauth2

type GrantType string

const (
	AuthorizationCode GrantType = "authorization_code"
	ClientCredentials GrantType = "client_credentials"
	Refreshing        GrantType = "refresh_token"
)

type ResponseType string

const (
	Code ResponseType = "code"
)

type TokenType string

const (
	Bearer TokenType = "bearer"
)
