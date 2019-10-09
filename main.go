package main

import (
	"gin-lab/routers"
	"gin-lab/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"os"
)

// @contact.name API Support
// @contact.url https://blog.weii.ink
// @contact.email wevsmy@gmail.com
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("$PORT must be set default %s", utils.Config.Host)
		port = utils.Config.Port
	}

	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + port)
}
