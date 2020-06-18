package server

import (
	"context"
	"log"
	pb "m_thor/proto"
	"m_thor/service"
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
