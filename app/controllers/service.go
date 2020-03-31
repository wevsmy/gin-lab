/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: service.go
@Time: 2019/10/08 下午5:43
*/

package controllers

import (
	"gin-lab/app/middleware"
	gin_jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @首页
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "hello world"
// @Router /hello/ [get]
func Index(c *gin.Context) {
	//返回结果
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "hello world",
	})
}

// @ping
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "v1 pong"
// @Router /ping [get]
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "v1 pong",
	})
}

// Method 测试业务逻辑接口
type MethodTestInterface interface {
	// CURD 方法具体实现接口
	//Query(ctx context.Context) ( error)
}

// @Name MethodTest
// @Description Method 方法测试API
type MethodTest struct {
	I MethodTestInterface
}

// @Summary get test
// @Description get string by ID
// @Tags method
// @Accept  json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} string "get 方法"
// @Failure 404 {string} string "not found"
// @Router /method/{id} [get]
func (m *MethodTest) GetOne(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "get one id:" + id,
	})
}

// @Summary get test
// @Description get ID
// @Tags method
// @Accept  json
// @Produce json
// @Param q query string false "name search by q" Format(email)
// @Success 200 {string} string "get list"
// @Failure 404 {string} string "not found"
// @Router /method [get]
func (m *MethodTest) GetList(c *gin.Context) {
	q := c.Request.URL.Query().Get("q")
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "get list id:" + q,
	})
}

type AddAccount struct {
	Name string `json:"name" example:"account name"`
}

// @Summary post test
// @Description post data
// @Tags method
// @Accept  json
// @Produce json
// @Param account body AddAccount true "Add account"
// @Success 200 {object} AddAccount
// @Failure 404 {string} string "not found"
// @Router /method [post]
func (m *MethodTest) Post(c *gin.Context) {
	var addAccount AddAccount
	if err := c.ShouldBindJSON(&addAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "post account:" + addAccount.Name,
	})
}

// @Summary put test
// @Description put data
// @Tags method
// @Accept  json
// @Produce json
// @Success 200 {string} string "put 方法"
// @Router /method [put]
func (m *MethodTest) Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "put 方法",
	})
}

// Account example
type Account struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

// @Summary delete test
// @Description Delete by account ID
// @Tags method
// @Accept  json
// @Produce json
// @Param  id path int true "Account ID" Format(int64)
// @Success 204 {object} Account
// @Router /method/{id} [delete]
func (m *MethodTest) Delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusNoContent, gin.H{
		"code":    http.StatusNoContent,
		"message": "delete id:" + id,
	})
}

type UpdateAccount struct {
	Name string `json:"name" example:"account name"`
}

// @Summary put test
// @Description Update a account
// @Tags method
// @Accept  json
// @Produce json
// @Param  id path int true "Account ID"
// @Param  account body UpdateAccount true "Update account"
// @Success 200 {object} Account
// @Failure 404 {string} string "not found"
// @Router /method/{id} [patch]
func (m *MethodTest) Patch(c *gin.Context) {
	id := c.Param("id")
	var updateAccount UpdateAccount
	if err := c.ShouldBindJSON(&updateAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "patch id:" + id + " new:" + updateAccount.Name,
	})
}

// @Summary head test
// @Description head data
// @Tags method
// @Accept  json
// @Produce json
// @Success 200 {string} string "head 方法"
// @Router /method [head]
func (m *MethodTest) Head(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "head 方法",
	})
}

// @Summary options test
// @Description options data
// @Tags method
// @Accept  json
// @Produce json
// @Success 200 {string} string "options 方法"
// @Router /method [options]
func (m *MethodTest) Options(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "options 方法",
	})
}

// @Summary login
// @Description post data
// @Accept  json
// @Produce json
// @Success 200 {string} string "login 方法"
// @Router /login [post]
func Login(c *gin.Context) {
	middleware.Auth_JWT.LoginHandler(c)
}

// @Summary logout
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "logout 方法"
// @Router /logout [get]
func Logout(c *gin.Context) {
	c.SetCookie(
		middleware.Auth_JWT.CookieName,
		"",
		int(middleware.Auth_JWT.TimeFunc().Add(middleware.Auth_JWT.Timeout).Unix()-time.Now().Unix()),
		"/",
		middleware.Auth_JWT.CookieDomain,
		middleware.Auth_JWT.SecureCookie,
		middleware.Auth_JWT.CookieHTTPOnly,
	)
	c.JSON(200, gin.H{"code": http.StatusOK, "message": "logout oj8k"})
}

func RefreshToken(c *gin.Context) {
	middleware.Auth_JWT.RefreshHandler(c)
}

func Hello(c *gin.Context) {
	claims := gin_jwt.ExtractClaims(c)
	user, _ := c.Get(middleware.IdentityKey)
	c.JSON(200, gin.H{
		"userID":    claims[middleware.IdentityKey],
		"userName":  user.(*middleware.User).UserName,
		"firstName": user.(*middleware.User).FirstName,
		"lastName":  user.(*middleware.User).LastName,
		"text":      "Hello World.",
	})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "test",
	})
}
