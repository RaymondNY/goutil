package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zqx.com/goutil/middlewares"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.Cors())
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello")
	})
	r.Run(":7777")
	//读取配置文件  viper可以监听配置的修改，还有个回调函数
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configfile")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(viper.Get("accound"))
	fmt.Println(viper.Get("detials.friends.info.tel"))
}
