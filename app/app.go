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
	"gin-lab/app/config"
	_ "gin-lab/app/models"
	"gin-lab/app/router"
	"github.com/gin-gonic/gin"
	"log"
	"os/user"
	"path/filepath"
)

// 全局结构体变量
type App struct {
	A config.Config
}

// app 应用
func Run() {
	u, _ := user.Current()
	filePath := filepath.Join(u.HomeDir, ".GinLabConfig", "config.yaml")
	c, e := config.New(filePath)
	if e != nil {
		log.Fatalln("e", e)
	}
	r := gin.New()
	router.Router(r)
	_ = r.Run(":" + c.Port)
}
