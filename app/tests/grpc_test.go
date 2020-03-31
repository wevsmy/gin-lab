//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: grpc_test.go
//@Time: 2020/3/31 下午3:03

package tests

import (
	"fmt"
	pb "gin-lab/app/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

func Test_grpc(t *testing.T) {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//初始化客户端
	c := pb.NewHelloClient(conn)

	//调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPCxxx"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.Message)

}
