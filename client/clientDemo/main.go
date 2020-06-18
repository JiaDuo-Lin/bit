package main

import (
	"context"
	thorPB "github.com/Kydaa/bit/service/m_thor/proto"
	userPB "github.com/Kydaa/bit/service/user/proto"
	"log"

	"google.golang.org/grpc"
)

const (
	thorAddress = "localhost:9090"
	userAddress = "localhost:9091"
)

func Token(token string) (status bool, err error) {
	// 建立grpc连接
	conn, err := grpc.Dial(thorAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect grpc:[%v]", err)
		return
	}

	defer conn.Close()
	// 新建客户端
	c := thorPB.NewServiceClient(conn)

	//testToken := "sajhgdshjkadkjaasdas"
	// 进行远程调用
	r, err := c.Token(context.TODO(), &thorPB.TokenRequest{Token:token})
	if err != nil {
		log.Fatalf("cant call remote grpc with err[%v]", err)
		return
	}
	log.Printf("Token result ", r.StatusCode)
	return r.StatusCode, nil
}

func login(id int64) string {
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect grpc:[%v]", err)
	}

	defer conn.Close()
	// 新建客户端
	c := userPB.NewServiceClient(conn)

	// 进行远程调用
	r, err := c.Login(context.TODO(), &userPB.LoginRequest{Id:id})
	if err != nil {
		log.Fatalf("cant call remote grpc with err[%v]", err)
	}
	log.Println("Token result ", r.Token, r.User)
	return r.Token
}

// RegisterGrpcDemoServer 注册grpc服务
func main() {
	token := login(12345)

	// 建立grpc连接
	Token(token)
}
