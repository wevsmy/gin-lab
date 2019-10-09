package main

import (
	"fmt"
	"gin-lab/docs"
	"gin-lab/routers"
	"gin-lab/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"os"
)

func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Gin-Lab API"
	docs.SwaggerInfo.Description = "This is a sample server gin-lab server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", utils.Config.Host, utils.Config.Port)
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("$PORT must be set default %s", utils.Config.Host)
		port = utils.Config.Port
	}

	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + port)
}
