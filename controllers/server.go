/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
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

// @method test
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "get 方法"
// @Router /method [get]
func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get 方法",
	})
}

// @method test
// @Description post data
// @Accept  json
// @Produce json
// @Success 200 {string} string "post 方法"
// @Router /method [post]
func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "post 方法",
	})
}

// @method test
// @Description put data
// @Accept  json
// @Produce json
// @Success 200 {string} string "put 方法"
// @Router /method [put]
func Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "put 方法",
	})
}

// @method test
// @Description delete data
// @Accept  json
// @Produce json
// @Success 200 {string} string "delete 方法"
// @Router /method [delete]
func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete 方法",
	})
}

// @method test
// @Description patch data
// @Accept  json
// @Produce json
// @Success 200 {string} string "patch 方法"
// @Router /method [patch]
func Patch(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "patch 方法",
	})
}

// @method test
// @Description head data
// @Accept  json
// @Produce json
// @Success 200 {string} string "head 方法"
// @Router /method [head]
func Head(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "head 方法",
	})
}

// @method test
// @Description options data
// @Accept  json
// @Produce json
// @Success 200 {string} string "options 方法"
// @Router /method [options]
func Options(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "options 方法",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}
