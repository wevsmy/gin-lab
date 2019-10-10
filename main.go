package main

import (
	"gin-lab/routers"
	"gin-lab/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// @contact.name API Support
// @contact.url https://blog.weii.ink
// @contact.email wevsmy@gmail.com
func main() {
	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + utils.Config.Port)
}
