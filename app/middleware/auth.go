/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: auth.go
@Time: 2019/10/10 下午5:54
*/

package middleware

import (
	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var Auth_JWT *gin_jwt.GinJWTMiddleware

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var IdentityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// 认证中间件初始化
func init() {
	var err error
	Auth_JWT, err = gin_jwt.New(&gin_jwt.GinJWTMiddleware{
		// Realm JWT标识
		Realm: "test zone",
		// Key 服务端密钥
		Key: []byte("secret key"),
		// Timeout token 过期时间
		Timeout: time.Minute * 50,
		// MaxRefresh token 更新时间  token有效期 = Timeout + MaxRefresh
		MaxRefresh: time.Minute * 10,
		// Authenticator 在登录接口中使用的验证方法，并返回验证成功后的用户对象。
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", gin_jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, gin_jwt.ErrFailedAuthentication
		},
		// Authorizator 登录后其他接口验证传入的 token 方法
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		// PayloadFunc 添加额外业务相关的信息
		PayloadFunc: func(data interface{}) gin_jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return gin_jwt.MapClaims{
					IdentityKey: v.UserName,
					"k":         "v",
				}
			}
			return gin_jwt.MapClaims{}
		},
		// Unauthorized 验证失败后设置错误信息
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// Login登录响应函数
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "login oj8k",
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
			})
		},
		// Refresh Token响应函数
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "refresh token oj8k",
				"token":   token,
				"expire":  expire.Format(time.RFC3339),
			})
		},
		// IdentityHandler 身份处理程序
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := gin_jwt.ExtractClaims(c)
			return &User{
				UserName:  claims[IdentityKey].(string),
				LastName:  "IdentityHandler 身份处理程序",
				FirstName: "业务相关" + claims["k"].(string),
			}
		},
		// IdentityKey 身份密钥
		IdentityKey: IdentityKey,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		// TokenLookup 设置 token 获取位置，一般默认在头部的 Authorization 中，或者 query的 token 字段，cookie 中的 jwt 字段。
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		// TokenHeadName Header中 token 的头部字段，默认常用名称 Bearer。
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		// TimeFunc 设置时间函数
		TimeFunc: time.Now,

		// return the token as a cookie
		// 将token作为cookie返回 便于开发测试
		SendCookie: true,
	})
	if err != nil {
		log.Fatalf("JWT Error: %v", err)
	}
}
