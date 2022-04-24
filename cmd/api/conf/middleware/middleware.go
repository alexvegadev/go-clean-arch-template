package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type userId struct{}

func PropagationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := "1000"

		ctx := context.WithValue(c.Request.Context(), userId{}, id)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
