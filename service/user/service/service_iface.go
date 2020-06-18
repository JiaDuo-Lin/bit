package service

import pb "github.com/Kydaa/bit/service/user/proto"
import (
	"context"
)

type IUserService interface {
	Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(context.Context, *pb.LoginRequest) (*pb.LoginResponse, error)
}
