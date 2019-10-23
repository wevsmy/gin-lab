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

package router

import (
	"gin-lab/app/controllers"
	_ "gin-lab/app/docs"
	"gin-lab/app/middleware"
	_ "gin-lab/app/utils"
	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

// 主路由入口(路由注册)
func Router(r *gin.Engine) {
	// 全局中间件
	r.Use(
		gin.Logger(),   // 日志记录
		gin.Recovery(), // 有panic时, 进行500的错误处理
		//middleware.TestMiddleware(), // 自定义测试中间件
	)
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// 首页
	r.GET("/", controllers.Index)
	// 静态文件
	staticRouter(r)
	// api
	v1ApiRouter(r)
	// 认证相关路由
	//authRouter(r)
	// 测试
	testRouter(r)
}

// v1 API路由入口
func v1ApiRouter(r *gin.Engine) {
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// websocket
	r.GET("/ws", controllers.HandleConnections)
	// upload
	upload := r.Group("/upload")
	{
		// 单文件上传
		upload.POST("/single", controllers.UploadSingle)
		upload.StaticFile("/single", "./app/static/single.html")
		// 多文件上传
		upload.POST("/multiple", controllers.UploadMultiple)
		upload.StaticFile("/multiple", "./app/static/multiple.html")
	}

	// 路由组
	v1 := r.Group("/v1")
	{
		// 二级路由
		v1.GET("/ping", controllers.Pong)

		method := v1.Group("/method")
		{
			var m *controllers.MethodTest
			method.GET(":id", m.GetOne)
			method.GET("", m.GetList)
			method.POST("", m.Post)
			method.PUT("", m.Put)
			method.DELETE(":id", m.Delete)
			method.PATCH(":id", m.Patch)
			method.HEAD("", m.Head)
			method.OPTIONS("", m.Options)
		}
	}
}

// 认证相关路由
func authRouter(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)

	r.NoRoute(middleware.Auth_JWT.MiddlewareFunc(), func(c *gin.Context) {
		claims := gin_jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", controllers.RefreshToken)
	auth.Use(middleware.Auth_JWT.MiddlewareFunc())
	{
		auth.GET("/hello", controllers.Hello)
	}
}

// 静态文件路由入口
func staticRouter(r *gin.Engine) {
	// 静态资源文件夹
	r.Static("/static", "./app/static")

	// 静态文件图标
	r.StaticFile("/favicon.ico", "./app/static/favicon.ico")
}

// 测试路由入口
func testRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		v, ex := c.Get("example")
		log.Println("牛逼：", v, ex)

		c.SetCookie("user_cookie", string("123456"), 1000, "/", "localhost", false, true)

		c.JSON(404, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("./app/templates/*.tmpl.html")
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
}
