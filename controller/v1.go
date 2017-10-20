package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func v1controller(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "login succeed",
			})
		})

		v1.POST("/logout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "logout succeed",
			})
		})

		user := v1.Group("/user")
		{
			user.GET("/info/:name", func(c *gin.Context) {
				name := c.Param("name")
				c.String(http.StatusOK, "Hello %s", name)
			})

			user.GET("/info/:name/*action", func(c *gin.Context) {
				name := c.Param("name")
				action := c.Param("action")
				message := name + " is " + action
				c.String(http.StatusOK, message)
			})

			user.GET("/welcome", func(c *gin.Context) {
				firstname := c.DefaultQuery("fistname", "Guest")
				lastname := c.Query("lastname")

				c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
			})

			user.POST("/form_post", func(c *gin.Context) {
				message := c.PostForm("message")
				nick := c.DefaultPostForm("nick", "anonymous")

				c.JSON(200, gin.H{
					"status":  "posted",
					"message": message,
					"nick":    nick,
				})
			})
		}
	}
}
