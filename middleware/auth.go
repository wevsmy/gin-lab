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
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(t *jwt.Token) (i interface{}, e error) {
				return []byte("SecretKey"), nil
			})

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
