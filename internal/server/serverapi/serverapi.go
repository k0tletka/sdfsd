package serverapi

import (
	"context"
	"fmt"
	pb "github.com/k0tletka/sdfsd/internal/protobuf"
	"github.com/k0tletka/sdfsd/internal/server/serverapi/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"

	"github.com/k0tletka/sdfsd/internal/server/config"
)

type ServerAPI struct {
	ctx  context.Context
	conf *config.ServerConfig
}

func (s *ServerAPI) Init(ctx context.Context, conf *config.ServerConfig) {
	s.ctx, s.conf = ctx, conf
}

func (s *ServerAPI) StartServer() error {
	listenAddr := fmt.Sprintf("%s:%d",
		s.conf.Config.ServerAPIConf.ListenAddr,
		s.conf.Config.ServerAPIConf.ListenPort,
	)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	if s.conf.Config.ServerAPIConf.UseSSL {
		creds, err := credentials.NewServerTLSFromFile(
			s.conf.Config.ServerAPIConf.CertFilePath,
			s.conf.Config.ServerAPIConf.KeyFilePath,
		)
		if err != nil {
			return err
		}

		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterServerAPIServer(grpcServer, handler.NewServerAPIHandler(s.ctx, s.conf))

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		errChan <- grpcServer.Serve(lis)
	}()

	select {
	case err := <-errChan:
		return err
	case <-s.ctx.Done():
		grpcServer.GracefulStop()
		<-errChan
		return nil
	}
}
