/*"""
@Version: V1.0
@Author: willson.wu
@License: Apache Licence
@Contact: willson.wu@goertek.com
@Site: goertek.com
@Software: GoLand
@File: main_test.go
@Time: 2019/10/08 下午5:45
*/

package tests

import (
	"fmt"
	"gin-lab/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()

	engine := gin.New()

	engine.GET("/ping", controllers.Pong)

	engine.ServeHTTP(rr, req)

	return rr
}

// 测试用例1
func TestServerBasic(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	fmt.Println(response.Body)
}
