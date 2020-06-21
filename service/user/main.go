package main

import (
	pb "github.com/Kydaa/bit/service/user/proto"
	"github.com/Kydaa/bit/service/user/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9091"
)

// RegisterJWTServer 注册微服务
func RegisterUserServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen with err:[%v]", err)
	}
	s := grpc.NewServer()
	ser := server.NewUserSever()
	pb.RegisterServiceServer(s, ser)
	//pb.RegisterGRPCServer(s)
	s.Serve(lis)
}

func main() {
	//fmt.Println("test")
	RegisterUserServer()
}
