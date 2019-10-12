/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: app.go.go
@Time: 2019/10/11 下午7:43
*/

package app

import (
	_ "gin-lab/app/models"
	"gin-lab/app/router"
	"github.com/gin-gonic/gin"
)

// 全局结构体变量
var App struct {
	Config *config
}

// app 应用
func Run() {
	App.Config.Init()
	r := gin.New()
	router.Router(r)
	_ = r.Run(":" + App.Config.Port)
}
