package main

import (
	pb "github.com/Kydaa/bit/service/m_thor/proto"
	"github.com/Kydaa/bit/service/m_thor/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9090"
)

// RegisterJWTServer 注册微服务
func RegisterJWTServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen with err:[%v]", err)
	}
	s := grpc.NewServer()
	ser := server.NewJWTServer()
	pb.RegisterServiceServer(s, ser)
	//pb.RegisterGRPCServer(s)
	s.Serve(lis)
}

func main() {
	//fmt.Println("test")
	RegisterJWTServer()
}
