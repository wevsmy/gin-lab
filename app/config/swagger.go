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

package config

import (
	"fmt"
	"gin-lab/app/docs"
)

// 初始化设置swagger的信息
func (c *config) swaggerInfoInit() {
	docs.SwaggerInfo.Title = "Swagger Gin-Lab API"
	docs.SwaggerInfo.Description = "This is a sample server gin-lab server."
	docs.SwaggerInfo.Version = "1.0"
	if c.Port == "80" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s", c.Host)
	} else {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", c.Host, c.Port)
	}
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
