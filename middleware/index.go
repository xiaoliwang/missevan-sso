package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func SendResponse() gin.HandlerFunc {
	return func (c *gin.Context) {
		start_time := time.Now()
		c.Next()
		latency := time.Since(start_time)
		c.Writer.Header().Set("X-forward-time", latency.String())
		if body, exists := c.Get("body"); true == exists {
			if code, exists := c.Get("code"); true == exists {
				c.SecureJSON(code.(int), body)
			} else {
				c.SecureJSON(http.StatusOK, body)
			}
		}
	}
}

func CheckAuth() gin.HandlerFunc {
	return func (c *gin.Context) {
		if sign := c.GetHeader("sign"); len(sign) > 0 {
			c.Next()
		} else {
			c.SecureJSON(http.StatusForbidden, "No Auth")
			c.Abort()
		}
	}
}
