package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
	"user-center/internal/consts"
	"user-center/internal/utils/jwtutils"
	"user-center/pkg/response"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if len(authorization) == 0 {
			response.AuthError(ctx, errors.New("登录已过期，请重新登录"))
			ctx.Abort()
			return
		}

		tokenStr := strings.Replace(authorization, "Bearer ", "", 1)
		claims, err := jwtutils.ParseJwt(tokenStr, consts.JwtSecret)
		if err != nil {
			response.AuthError(ctx, errors.New("token异常，请重新登录"))
			ctx.Abort()
			return
		}
		ctx.Set(consts.ContextKeyUser, claims.User)
		ctx.Set(consts.ContextKeyToken, tokenStr)
		ctx.Next()
	}
}
