package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Test() gin.HandlerFunc {
	return func (c *gin.Context) {
		start_time := time.Now()
		c.Set("example", "12345")

		c.Next()

		latercy := time.Since(start_time)
		c.Writer.Header().Set("X-forward-time", latercy.String())
		print("duration", latercy.String())

	}
}
