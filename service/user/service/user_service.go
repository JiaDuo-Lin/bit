package service

import (
	"github.com/pkg/errors"
	"user/model"
	pb "user/proto"
)
import (
	"context"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Register(ctx context.Context,
	req *pb.RegisterRequest) (resp *pb.RegisterResponse, err error) {

	user := model.NewUser(req.User.Id, req.User.Name, req.User.Tags)
	user.Add()
	resp = &pb.RegisterResponse{StatusCode: true}
	return
}

func (us *UserService) Login(ctx context.Context,
	req *pb.LoginRequest) (resp *pb.LoginResponse, err error) {

	if req.Id == 12345 {
		userResp := &pb.User{Name: "Mike", Id: 12345, Tags: []string{"数据结构", "后端"}}
		token := "testToken"
		resp = &pb.LoginResponse{StatusCode: true, User: userResp, Token: token}
		return
	}

	user := model.NewUser(req.Id, "", nil)
	if !user.IsRegistered() {
		return writeErr("User is not registered")
	}

	err = user.Load()
	if err != nil {
		return writeErr("Error in load user data")
	}

	token, err := model.CreateToken(user.Name, user.ID)
	if err != nil {
		return writeErr("Error in create token")
	}

	userResp := &pb.User{Name: user.Name, Id: user.ID, Tags: user.Tags}
	resp = &pb.LoginResponse{StatusCode: true, User: userResp, Token: token.Token}
	return
}

func writeErr(reason string) (resp *pb.LoginResponse, err error) {
	err = errors.Wrapf(err, reason)
	resp = &pb.LoginResponse{StatusCode: false}
	return
}
