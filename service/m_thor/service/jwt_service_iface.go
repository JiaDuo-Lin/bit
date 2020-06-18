package service

import (
	"context"
	pb "m_thor/proto"
)

type IJWTService interface {
	Token(context.Context, *pb.TokenRequest) (*pb.TokenResponse, error)
}
