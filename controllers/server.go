/*"""
@Version: V1.0
@Author: willson.wu
@License: Apache Licence
@Contact: willson.wu@goertek.com
@Site: goertek.com
@Software: GoLand
@File: server.go
@Time: 2019/10/08 下午5:43
*/

package controllers

import (
	"github.com/gin-gonic/gin"
)

// @首页
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "hello world"
// @Router /hello/ [get]
func Index(c *gin.Context) {
	//返回结果
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

// @ping
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "v1 pong"
// @Router /ping [get]
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "v1 pong",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
