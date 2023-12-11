package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

const (
	XRequestIDKey = "X-Request-ID"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader(XRequestIDKey)

		if rid == "" {
			rid = strings.ReplaceAll(uuid.NewV4().String(), "-", "")
			ctx.Request.Header.Set(XRequestIDKey, rid)
			ctx.Set(XRequestIDKey, rid)
		}

		// Set XRequestIDKey header
		ctx.Writer.Header().Set(XRequestIDKey, rid)
		ctx.Next()
	}
}
