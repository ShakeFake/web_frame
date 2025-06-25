package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"wilikidi/gin/controller"
	"wilikidi/gin/utils"
)

// @title WebFrame 接口文档
// @version 2.0
// @description

// @contact.name
// @contact.url
// @contact.email

// @license.name MIT
// @license.url https://github.com/xinliangnote/go-gin-api/blob/master/LICENSE

// @securityDefinitions.apikey  LoginToken
// @in                          header
// @name                        token

// @BasePath /
func main() {
	utils.InitLogger()

	utils.Config()

	// 初始化一个 validator
	binding.Validator = utils.BindValidator()

	// mysql 相关使用。

	// gin 框架注册制中间件
	router := gin.Default()
	// set the cors。when https
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,accept, content-type, referer, sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, Tagvalue, User-Agent")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	})

	router.GET("/ping", controller.Version)
	router.POST("/validator", controller.Validator)

	//router.RunTLS(":5002", "server.crt", "server.key")
	router.Run(":8091")
}
