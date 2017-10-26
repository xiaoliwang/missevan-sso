package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func BaseController(r *gin.Engine) {

	// 访问 /ping 查看服务是否运行
	r.GET("/ping", testPing(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		c.Writer.Header().Set("X-forward-time2", "1234")
	})

	// v1 版本接口
	apiController(r)
}

func testPing() gin.HandlerFunc {
	return func(c *gin.Context) {
		start_time := time.Now()
		c.Set("example", "12345")
		c.Writer.Header().Set("X-forward-time1", "1234")
		c.Next()

		latercy := time.Since(start_time)

		print("duration", latercy.String())
	}
}