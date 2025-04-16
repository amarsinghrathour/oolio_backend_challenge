package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("api_key")
		fmt.Println("api key:", key)
		if key != apiKey {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		c.Next()
	}
}
