/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: logrus.go
@Time: 2019/10/10 下午13:41
*/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// log's日志记录中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 测试自定义中间件
func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		// 在gin上下文中定义变量
		c.Set("example", "12345")
		// 开始时间
		startTime := time.Now()

		//处理请求
		c.Next()

		// 请求后
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUrl := c.Request.RequestURI
		// 请求IP
		clientIP := c.ClientIP()
		// 状态码
		statusCode := c.Writer.Status()
		// 错误信息
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		outString := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s  %s %s %s",
			time.Now().Format("2006/01/02 - 15:04:05"),
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUrl,
			errorMessage,
			"TestMiddleware",
		)
		fmt.Println(outString)
	}
}
