package service

import (
	"github.com/Kydaa/bit/service/m_thor/model"
	pb "github.com/Kydaa/bit/service/m_thor/proto"
)
import "context"

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (j *JWTService) Token(ctx context.Context,
	req *pb.TokenRequest) (*pb.TokenResponse, error) {

	status := model.CheckToken(req.Token)
	resp := &pb.TokenResponse{StatusCode: status}
	return resp, nil
}
