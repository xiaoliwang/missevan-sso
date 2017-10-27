package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
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
		sign := c.GetHeader("sign")
		if len(sign) <= 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, "No Auth")
		}

		row, error := c.GetRawData()
		if nil != error {
			r.
		}



		if sign := c.GetHeader("sign"); len(sign) > 0 {
			c.Next()
		} else {

		}
	}
}

var mac = hmac.New(sha256.New, []byte("1234"))

func checkAuth(sign []byte, c *gin.Context) bool {
	if row, error := c.GetRawData(); nil == error {
		mac.Write(row)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(sign, expectedMAC)
	} else {
		return false
	}
}