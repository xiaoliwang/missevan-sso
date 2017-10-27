package main

import (
	"github.com/gin-gonic/gin"
	"missevan_sso/middleware"
	"missevan_sso/controller"
)

func main() {
	// router := gin.Default()
	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.SendResponse(),

	)

	controller.BaseController(router)

	router.Run(":3000")
}
