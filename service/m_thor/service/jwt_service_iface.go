/*
	鉴权微服务借口
	author：Kyda
	date:06.18.2020
*/

package service

import (
	"context"
	pb "github.com/Kydaa/bit/service/m_thor/proto"
)

type IJWTService interface {
	Token(context.Context, *pb.TokenRequest) (*pb.TokenResponse, error)
}
