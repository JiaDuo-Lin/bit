//+build ignore
package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "user/proto"
	"user/server"
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
