//@Version: V1.0
//@Author: wevsmy
//@License: Apache Licence
//@Contact: wevsmy@gmail.com
//@Site: blog.weii.ink
//@Software: GoLand
//@File: grpc.go
//@Time: 2020/3/31 下午1:34

package server

import (
	pb "gin-lab/app/protos"
	"golang.org/x/net/context"
)

//定义一个helloServer并实现约定的接口
type HelloService struct {
}

func (h HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello-" + in.Name + "-end"
	return resp, nil
}

//var HelloServer = helloService{}
