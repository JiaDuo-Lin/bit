package main

import (
	"context"
	pb "github.com/Kydaa/bit/service/m_thor/proto"
	"log"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

// RegisterGrpcDemoServer 注册grpc服务
func main() {
	// 建立grpc连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect grpc:[%v]", err)
	}

	defer conn.Close()
	// 新建客户端
	c := pb.NewServiceClient(conn)

	testToken := "sajhgdshjkadkjaasdas"
	// 进行远程调用
	r, err := c.Token(context.TODO(), &pb.TokenRequest{Token:testToken})
	if err != nil {
		log.Fatalf("cant call remote grpc with err[%v]", err)
	}
	log.Printf("Token result ", r.StatusCode)
}
