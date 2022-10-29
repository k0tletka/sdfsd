package handler

import (
	"context"
	"fmt"
	srvconfig "github.com/k0tletka/sdfsd/internal/config"
	"github.com/k0tletka/sdfsd/internal/fs"
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

func (s *ServerAPIHandler) GetPools(ctx context.Context, _ *emptypb.Empty) (*pb.PoolListResponse, error) {
	pools := s.conf.VolDispatcher.GetPools()
	pbPools := make([]*pb.Pool, 0, len(pools))

	for _, pool := range pools {
		pbPools = append(pbPools, &pb.Pool{
			PoolName: pool.Name,
			PoolMode: pb.PoolMode(pool.WorkMode),
		})
	}

	return &pb.PoolListResponse{
		PoolList: pbPools,
	}, nil
}

func (s *ServerAPIHandler) GetPoolInfo(ctx context.Context, req *pb.PoolInfoRequest) (*pb.PoolInfoResponse, error) {
	pool, err := s.conf.VolDispatcher.GetPool(req.PoolName)
	if err == nil {
		return &pb.PoolInfoResponse{
			Pool: &pb.Pool{
				PoolName: pool.Name,
				PoolMode: pb.PoolMode(pool.WorkMode),
			},
		}, nil
	}

	fserr, ok := err.(fs.FSError)
	if ok {
		if fserr == fs.ErrNotFound {
			return &pb.PoolInfoResponse{
				ErrorInfo: &pb.ErrorInfo{
					ErrorCode:    pb.ErrorCode_OBJECT_NOT_FOUND,
					ErrorMessage: fserr.Error(),
				},
			}, nil
		}

		return nil, fmt.Errorf("FS error occured: %w", fserr)
	}

	return nil, err
}

func (s *ServerAPIHandler) GetVolumes(ctx context.Context, _ *emptypb.Empty) (*pb.VolumeListResponse, error) {
	volumes := s.conf.VolDispatcher.GetVolumes()
	pbVolumes := make([]*pb.Volume, 0, len(volumes))

	for _, volume := range volumes {
		if volume.Pool == "" {
			continue
		}

		pbVolumes = append(pbVolumes, &pb.Volume{
			VolumeName: volume.Name,
			VolumeSize: volume.VolumeSize,
			PoolName:   volume.Pool,
		})
	}

	return &pb.VolumeListResponse{
		VolumeList: pbVolumes,
	}, nil
}
