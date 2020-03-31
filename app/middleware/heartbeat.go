//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: heartbeat.go
//@Time: 2020/3/31 下午6:20

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	// DefaultHeaderName default header name
	DefaultHeaderName = "X-Health-Check"

	// DefaultHeaderValue default header value
	DefaultHeaderValue = "1"

	// DefaultResponseCode default response code
	DefaultResponseCode = http.StatusOK

	// DefaultResponseText default response text
	DefaultResponseText = "ok"

	// DefaultConfig default config
	DefaultConfig = Config{
		HeaderName:   DefaultHeaderName,
		HeaderValue:  DefaultHeaderValue,
		ResponseCode: DefaultResponseCode,
		ResponseText: DefaultResponseText}
)

// Config holds the configuration values
type Config struct {
	HeaderName   string
	HeaderValue  string
	ResponseCode int
	ResponseText string
}

// Default creates a new middileware with the default configuration
// 心跳检测中间件
func Heartbeat() gin.HandlerFunc {
	return New(DefaultConfig)
}

// New creates a new middileware with the `cfg`
func New(cfg Config) gin.HandlerFunc {
	if cfg.HeaderName == "" {
		cfg.HeaderName = DefaultHeaderName
	}
	if cfg.HeaderValue == "" {
		cfg.HeaderValue = DefaultHeaderValue
	}
	if cfg.ResponseCode == 0 {
		cfg.ResponseCode = DefaultResponseCode
	}

	return func(ctx *gin.Context) {
		if ctx.GetHeader(cfg.HeaderName) == cfg.HeaderValue {
			ctx.String(cfg.ResponseCode, cfg.ResponseText)
			ctx.Abort()
		}

		ctx.Next()
	}
}

// https://github.com/RaMin0/gin-health-check

//func main() {
//	router := gin.Default()
//	router.Use(healthcheck.New(healthcheck.Config{
//		HeaderName:   "X-Custom-Header",
//		HeaderValue:  "customValue",
//		ResponseCode: http.StatusTeapot,
//		ResponseText: "teapot",
//	}))
//}
//$ curl -iL -XGET -H "X-Custom-Header: customValue" http://localhost
//  # HTTP/1.1 418 I'm a teapot
//  # Content-Length: 6
//  # Content-Type: text/plain; charset=utf-8
//  #
//  # teapot
