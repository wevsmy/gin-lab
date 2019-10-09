/*"""
@Version: V1.0
@Author: willson.wu
@License: Apache Licence
@Contact: willson.wu@goertek.com
@Site: goertek.com
@Software: GoLand
@File: router.go
@Time: 2019/10/08 下午6:13
*/

package routers

import (
	"gin-lab/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiRouter(r *gin.Engine) {
	// 首页
	r.GET("/", controllers.Index)
	// 路由组
	v1 := r.Group("v1")
	{
		// 二级路由
		v1.GET("/ping", controllers.Pong)
	}
}

func staticRouter(r *gin.Engine) {

	// 静态资源文件夹
	r.Static("/statics", "statics")

	// 静态文件图标
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
}

// 主路由入口
func Router(r *gin.Engine) {

	// api
	apiRouter(r)
	// 静态文件
	staticRouter(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong123",
		})
	})

	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.tmpl.html")
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	//r.GET("/cookie", func(c *gin.Context) {
	//
	//	cookie, err := c.Cookie("gin_cookie")
	//
	//	if err != nil {
	//		cookie = "NotSet"
	//		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	//		fmt.Println("asdasdasd")
	//	}
	//
	//	fmt.Printf("Cookie value: %s \n", cookie)
	//})
}
