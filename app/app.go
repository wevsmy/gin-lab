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
	"gin-lab/app/routers"
	"gin-lab/app/utils"
	"github.com/gin-gonic/gin"
)

// app 应用初始化
func Init() {
	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + utils.Config.Port)
}
