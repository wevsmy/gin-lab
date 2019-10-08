package main

import (
	"log"
	"os"

	"gin-lab/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("$PORT must be set default 8080")
		port = "8080"
	}

	r := gin.New()
	routers.Router(r)
	_ = r.Run(":" + port)
}
