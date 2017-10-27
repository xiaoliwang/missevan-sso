package controller

import (
	"github.com/gin-gonic/gin"
)

func BaseController(r *gin.Engine) {

	// 访问 /ping 查看服务是否运行
	r.GET("/ping", func(c *gin.Context) {
		c.Set("body", "pong")
	})

	// v1 版本接口
	apiController(r)
}