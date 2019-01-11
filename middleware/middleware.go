package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware add request id to each request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Request-Id", uuid.New().String())
		ctx.Next()
	}
}
