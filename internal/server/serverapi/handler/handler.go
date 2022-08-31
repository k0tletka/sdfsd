package handler

import (
	"context"

	srvconfig "github.com/k0tletka/sdfsd/internal/config"
	pb "github.com/k0tletka/sdfsd/internal/protobuf"
	"github.com/k0tletka/sdfsd/internal/server/config"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAPIHandler struct {
	ctx  context.Context
	conf *config.ServerConfig

	pb.UnimplementedServerAPIServer
}

func NewServerAPIHandler(ctx context.Context, conf *config.ServerConfig) *ServerAPIHandler {
	return &ServerAPIHandler{
		ctx:  ctx,
		conf: conf,
	}
}

func (s *ServerAPIHandler) GetServerInfo(ctx context.Context, _ *emptypb.Empty) (*pb.ServerInfoResponse, error) {
	return &pb.ServerInfoResponse{
		ServerName: s.conf.Config.ServerName,
		ApiVersion: srvconfig.ApiVersion,
	}, nil
}
