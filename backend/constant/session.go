package constant

const (
	AuthMethodSession = "session"
	SessionName       = "psession"

	AuthMethodJWT = "jwt"
	JWTHeaderName = "Authorization"
	JWTLegacyHeaderName = "PanelAuthorization"
	JWTBufferTime = 259200
	JWTIssuer     = "1Panel"

	PasswordExpiredName = "expired"
)
