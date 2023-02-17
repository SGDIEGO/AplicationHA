package middleware

import "github.com/gin-gonic/gin"

type Logger struct{}

func NewMiddlewareLogger() *Logger {
	return &Logger{}
}

func (log *Logger) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userPass", "diego")

		ctx.Next()
	}
}
