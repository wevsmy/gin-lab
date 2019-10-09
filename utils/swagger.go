/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: swagger.go
@Time: 2019/10/9 下午5:25
*/

package utils

import (
	"fmt"
	"gin-lab/docs"
)

func init() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Gin-Lab API"
	docs.SwaggerInfo.Description = "This is a sample server gin-lab server."
	docs.SwaggerInfo.Version = "1.0"
	if Config.Port == "80" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s", Config.Host)
	} else {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", Config.Host, Config.Port)
	}
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
