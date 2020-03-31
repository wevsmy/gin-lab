/*"""
@Version: V1.0
@Author: wevsmy
@License: Apache Licence
@Contact: wevsmy@gmail.com
@Site: blog.weii.ink
@Software: GoLand
@File: app.go.go
@Time: 2019/10/11 下午7:43
*/

package app

import (
	"gin-lab/app/config"
	_ "gin-lab/app/models"
	pb "gin-lab/app/protos"
	"gin-lab/app/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os/user"
	"path/filepath"
	"strings"
)

// 全局结构体变量
type App struct {
	A config.Config
}

// app 应用
func Run() {
	// 读物配置文件信息
	u, _ := user.Current()
	filePath := filepath.Join(u.HomeDir, ".GinLabConfig", "config.yaml")
	c, e := config.New(filePath)
	if e != nil {
		log.Fatalln("e", e)
	}

	// 创建主侦听器。
	l, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Fatal(err)
	}

	// 创建复用链接
	m := cmux.New(l)

	// 按顺序匹配连接：
	// 首先使用grpc，然后使用HTTP，否则使用RPC / TCP。
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())
	// 表示尚未匹配的任何内容。
	trpcL := m.Match(cmux.Any())

	// 创建协议服务器
	grpcS := grpc.NewServer()
	s := new(server.HelloService)

	pb.RegisterHelloServer(grpcS, s)

	httpS := &http.Server{
		Handler: server.GinServer(),
	}

	trpcS := rpc.NewServer()
	//trpcS.Register(&ExampleRPCRcvr{})

	// 对服务器使用多路侦听器。
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)
	go trpcS.Accept(trpcL)

	// 开始服务
	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}
}
