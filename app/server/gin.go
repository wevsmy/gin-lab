//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: gin.go
//@Time: 2020/3/31 下午2:49

package server

import (
	"gin-lab/app/router"
	_ "gin-lab/app/router"
	"github.com/gin-gonic/gin"
)

func GinServer() *gin.Engine {
	engine := gin.New()
	router.Router(engine)
	return engine
}
