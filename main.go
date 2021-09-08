package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zqx.com/goutil/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})
	r.Run(":7777")
}
