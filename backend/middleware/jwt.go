// 此文件在GPL-3.0协议下开源
// 修改者：bobwu 2026-09-01
package middleware

import (
	"github.com/1Panel-dev/1Panel/backend/app/api/v1/helper"
	"github.com/1Panel-dev/1Panel/backend/constant"
	jwtUtils "github.com/1Panel-dev/1Panel/backend/utils/jwt"

	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.GetHeader(constant.JWTHeaderName)
		if rawToken == "" {
			rawToken = c.GetHeader(constant.JWTLegacyHeaderName)
		}
		if rawToken == "" {
			rawToken = c.Query("token")
		}
		token := strings.TrimSpace(rawToken)
		if strings.HasPrefix(strings.ToLower(token), "bearer ") {
			token = strings.TrimSpace(token[7:])
		}
		if token == "" {
			helper.ErrorWithDetail(c, constant.CodeErrUnauthorized, constant.ErrTypeNotLogin, nil)
			return
		}
		j := jwtUtils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			helper.ErrorWithDetail(c, constant.CodeErrUnauthorized, constant.ErrTypeNotLogin, err)
			return
		}
		if claims.Name == "" && claims.Subject != "" {
			claims.Name = claims.Subject
		}
		c.Set("claims", claims)
		c.Set("authMethod", constant.AuthMethodJWT)
		c.Next()
	}
}
