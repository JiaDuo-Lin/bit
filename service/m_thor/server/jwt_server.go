package server

import (
	"context"
	pb "github.com/Kydaa/bit/service/m_thor/proto"
	"github.com/Kydaa/bit/service/m_thor/service"
	"log"
)

type JWTServer struct {
	jwtService service.IJWTService
}

func NewJWTServer() *JWTServer {
	return &JWTServer{jwtService: service.NewJWTService()}
}

func (js *JWTServer) Token(ctx context.Context,
	req *pb.TokenRequest) (resp *pb.TokenResponse, err error) {
	resp, err = js.jwtService.Token(ctx, req)
	if err != nil {
		log.Fatalln("Err in Token server:", err)
	}
	return
}
