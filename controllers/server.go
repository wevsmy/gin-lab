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

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "v1 pong",
	})
}
