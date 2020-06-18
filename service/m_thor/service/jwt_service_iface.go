package service

import (
	"context"
	pb "github.com/Kydaa/bit/service/m_thor/proto"
)

type IJWTService interface {
	Token(context.Context, *pb.TokenRequest) (*pb.TokenResponse, error)
}
