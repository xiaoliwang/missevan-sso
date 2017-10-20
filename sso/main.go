package main

import (
	"github.com/gin-gonic/gin"
	"missevan_sso/controller"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	controller.BaseController(r)

	r.Run(":3000")
}
