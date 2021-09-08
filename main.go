package main

import (
	"github.com/gin-gonic/gin"
	"zqx.com/goutil/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
}
