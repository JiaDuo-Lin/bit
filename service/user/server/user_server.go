package server

import (
	"log"
	pb "user/proto"
	"user/service"
)
import "context"

type UserServer struct {
	userService service.IUserService
}

func NewUserSever() *UserServer {
	return &UserServer{
		userService: service.NewUserService(),
	}
}

func (us *UserServer) Register(ctx context.Context,
	req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {

	resp, err = us.userService.Register(ctx, req)
	if err != nil {
		log.Fatalln("Error in register server:", err)
	}
	return
}

func (us *UserServer) Login(ctx context.Context,
	req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {

	resp, err = us.userService.Login(ctx, req)
	if err != nil {
		log.Fatalln("Error in login server:", err)
	}
	return
}
