package main

import (
	"gin-lab/docs"
	"gin-lab/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"os"
)

func main() {

	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Gin-Lab API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set default 8080")
		port = "8080"
	}

	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + port)

}
