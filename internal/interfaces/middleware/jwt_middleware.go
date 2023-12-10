package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"user-center/internal/infrastructure/consts"
	"user-center/internal/infrastructure/utils/jwtutils"
	"user-center/pkg/code"
	"user-center/pkg/errors"
	"user-center/pkg/response"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if len(authorization) == 0 {
			response.Error(ctx, errors.WithCode(code.ErrMissingToken, "authorization为空"))
			ctx.Abort()
			return
		}

		tokenStr := strings.Replace(authorization, "Bearer ", "", 1)
		claims, err := jwtutils.ParseJwt(tokenStr, consts.JwtSecret)
		if err != nil {
			response.Error(ctx, errors.WithCode(code.ErrTokenInvalid, "token解析异常，请重新登录"))
			ctx.Abort()
			return
		}
		ctx.Set(consts.ContextKeyUser, claims.User)
		ctx.Set(consts.ContextKeyToken, tokenStr)
		ctx.Next()
	}
}
