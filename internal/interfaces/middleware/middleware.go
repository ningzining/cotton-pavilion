package middleware

import "github.com/gin-gonic/gin"

func DefaultMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		gin.Recovery(),
		Cors(),
		RequestID(),
		Logger(),
	}
}
