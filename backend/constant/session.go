// 此文件在GPL-3.0协议下开源
// 修改者：bobwu 2026-09-01
package constant

const (
	AuthMethodSession = "session"
	SessionName       = "psession"

	AuthMethodJWT       = "jwt"
	JWTHeaderName       = "Authorization"
	JWTLegacyHeaderName = "PanelAuthorization"
	JWTBufferTime       = 259200
	JWTIssuer           = "1Panel"

	PasswordExpiredName = "expired"
)
