package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/consts"
	"github.com/ningzining/cotton-pavilion/internal/infrastructure/util/jwtutil"
	"github.com/ningzining/cotton-pavilion/pkg/code"
	"github.com/ningzining/cotton-pavilion/pkg/errors"
	"github.com/ningzining/cotton-pavilion/pkg/response"
	"strings"
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
		claims, err := jwtutil.ParseJwt(tokenStr, consts.JwtSecret)
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
