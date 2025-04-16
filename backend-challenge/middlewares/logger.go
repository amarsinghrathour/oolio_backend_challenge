package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("%s %s | %d | %v\n", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	}
}
