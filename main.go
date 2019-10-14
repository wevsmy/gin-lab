package main

import (
	"gin-lab/app"
	_ "github.com/heroku/x/hmetrics/onload"
)

// @contact.name API Support
// @contact.url https://blog.weii.ink
// @contact.email wevsmy@gmail.com
func main() {
	//app.Run()
	app.TestAliPay()
}
