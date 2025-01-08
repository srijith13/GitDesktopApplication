package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("request-id", uuid.New())
		c.Next()
	}
}
