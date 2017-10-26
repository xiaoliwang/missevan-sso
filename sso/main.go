package main

import (
	"github.com/gin-gonic/gin"
	"missevan_sso/controller"
	_ "missevan_sso/middleware"
)

func main() {
	// gin.Default == gin.New().Use(gin.Logger(), gin.Recovery())
	router := gin.Default()

	/*r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Test())*/

	controller.BaseController(router)

	router.Run(":3000")
}
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

