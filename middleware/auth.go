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
	"time"
)

var Auth_JWT *gin_jwt.GinJWTMiddleware

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func HelloHandler(c *gin.Context) {
	claims := gin_jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

func init() {
	var err error
	Auth_JWT, err = gin_jwt.New(&gin_jwt.GinJWTMiddleware{
		// Realm JWT标识
		Realm: "test zone",
		// Key 服务端密钥
		Key: []byte("secret key"),
		// Timeout token 过期时间
		Timeout: time.Hour,
		// MaxRefresh token 更新时间
		MaxRefresh: time.Hour,
		// IdentityKey 身份密钥
		IdentityKey: identityKey,
		// PayloadFunc 添加额外业务相关的信息
		PayloadFunc: func(data interface{}) gin_jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return gin_jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return gin_jwt.MapClaims{}
		},
		// IdentityHandler 身份处理程序
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := gin_jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
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
		// Unauthorized 验证失败后设置错误信息
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
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
	})
	if err != nil {
		log.Fatalf("JWT Error: %v", err)
	}
}

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
		//	func(t *jwt.Token) (i interface{}, e error) {
		//		return []byte("SecretKey"), nil
		//	})

		//fmt.Println("auth:", token, err)

		//a := &jwt.GinJWTMiddleware{
		//	Realm:         "gin jwt",
		//	Key:           []byte("secret key"),
		//	Timeout:       time.Hour,
		//	MaxRefresh:    time.Hour,
		//	PayloadFunc:   func(data interface{}) jwt.MapClaims {},
		//	Authenticator: func(c *gin.Context) (interface{}, error) {},
		//	Authorizator:  func(data interface{}, c *gin.Context) bool {},
		//	Unauthorized:  func(c *gin.Context, code int, message string) {},
		//	TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		//	// TokenLookup: "query:token",
		//	// TokenLookup: "cookie:token",
		//	TokenHeadName: "Bearer",
		//	TimeFunc:      time.Now,
		//}

		//if err == nil {
		//	if token.Valid {
		//		c.Next()
		//	} else {
		//		c.String(http.StatusUnauthorized, "Token is not valid")
		//	}
		//} else {
		//	c.String(http.StatusUnauthorized, "Unauthorized access to this resource")
		//}

	}
}
