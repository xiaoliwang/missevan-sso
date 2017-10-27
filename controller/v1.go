package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"missevan_sso/model"
	"missevan_sso/middleware"
)

func apiController(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.CheckAuth())
	{
		v1 := api.Group("/v1")
		{
			// 登录接口
			v1.POST("/loginJson", func(c *gin.Context) {
				var form model.LoginForm

				if err := c.ShouldBindJSON(&form); err == nil {
					if form.Account == "jiepengthegreat@126.com" &&
						form.Password == "cao123456" {
						names := []string{"lena", "austin", "foo"}
						c.Set("body", names)
						c.Writer.Header().Set("fuck", "haha")
					} else {
						c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
					}
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
			})

			v1.POST("/loginForm", func(c *gin.Context) {
				var form model.LoginForm

				if err := c.ShouldBind(&form); err == nil {
					if form.Account == "jiepengthegreat@126.com" &&
						form.Password == "cao123456" {
						c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
					} else {
						c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
					}
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				}
			})

			// 登出接口
			v1.POST("/logout", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"success": true,
					"message": "logout succeed",
				})
			})
		}
	}
}
