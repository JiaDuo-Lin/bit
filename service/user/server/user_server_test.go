package server

import (
	"context"
	pb "demo/proto"
	"github.com/pkg/errors"
	"testing"
)

func TestUserServer(t *testing.T) {
	s := NewUserSever()
	resp, err := s.Login(context.Background(), &pb.LoginRequest{Id: 12345})

	if err != nil {
		err = errors.Wrapf(err, "Test has err[%v] with resp[%s]", err, resp)
	}

	//t.Fatal(resp.User, resp.Token)

}
