/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: router.go
@Time: 2019/10/08 下午6:13
*/

package routers

import (
	"gin-lab/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

// 主路由入口(路由注册)
func Router(r *gin.Engine) {
	// 首页
	r.GET("/", controllers.Index)
	// 静态文件
	staticRouter(r)
	// api
	v1ApiRouter(r)
	// 测试
	testRouter(r)
}

// v1 API路由入口
func v1ApiRouter(r *gin.Engine) {
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由组
	v1 := r.Group("v1")
	{
		// 二级路由
		v1.GET("/ping", controllers.Pong)

		method := v1.Group("/method")
		{
			method.GET(":id", controllers.Get)
			method.GET("", controllers.Get)
			method.POST("", controllers.Post)
			method.PUT("", controllers.Put)
			method.DELETE("", controllers.Delete)
			method.PATCH("", controllers.Patch)
			method.HEAD("", controllers.Head)
			method.OPTIONS("", controllers.Options)
		}
	}
}

// 静态文件路由入口
func staticRouter(r *gin.Engine) {
	// 静态资源文件夹
	r.Static("/statics", "statics")

	// 静态文件图标
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
}

// 测试路由入口
func testRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
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
