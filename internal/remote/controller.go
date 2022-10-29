package remote

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/k0tletka/sdfsd/internal/fs/enum"
	pb "github.com/k0tletka/sdfsd/internal/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"time"
)

var (
	RemoteController *RemoteServerController
)

const (
	remoteServersConfigLocation = "/var/sdfds/remoteServers.json"

	grpcRequestTimeout = 10 * time.Second
)

type RemoteServer struct {
	ConnectionString string
	UseSSL           bool
	ServerName       string

	ctx             context.Context
	serverApiClient pb.ServerAPIClient
	isInitialized   bool
	remoteVolumes   []RemoteVolume
	remotePools     []RemotePool
}

func (r *RemoteServer) IsInitialized() bool {
	return r.isInitialized
}

func (r *RemoteServer) InitializeRemoteServer(ctx context.Context) error {
	var opts []grpc.DialOption

	if r.UseSSL {
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{}),
		))
	}

	r.ctx = ctx
	ctx, cancel := context.WithTimeout(ctx, grpcRequestTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, r.ConnectionString, opts...)
	if err != nil {
		return err
	}

	r.serverApiClient = pb.NewServerAPIClient(conn)

	if err := r.LoadRemotePools(); err != nil {
		return err
	}
	return r.LoadRemoteVolumes()
}

func (r *RemoteServer) LoadRemoteVolumes() error {
	if !r.IsInitialized() {
		return ErrRemoteServerNotInitialized
	}

	// TODO: make request to get all volumes from remote server
	ctx, cancel := context.WithTimeout(r.ctx, grpcRequestTimeout)
	defer cancel()

	remoteVolumeResp, err := r.serverApiClient.GetVolumes(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	r.remoteVolumes = make([]RemoteVolume, 0, len(remoteVolumeResp.VolumeList))

	for _, pbVolume := range remoteVolumeResp.VolumeList {
		r.remoteVolumes = append(r.remoteVolumes, RemoteVolume{
			Name: pbVolume.VolumeName,
			Size: pbVolume.VolumeSize,
			Pool: pbVolume.PoolName,
		})
	}

	return nil
}

func (r *RemoteServer) LoadRemotePools() error {
	if !r.IsInitialized() {
		return ErrRemoteServerNotInitialized
	}

	// TODO: make request to get all pools from remote server
	ctx, cancel := context.WithTimeout(r.ctx, grpcRequestTimeout)
	defer cancel()

	remotePoolsResp, err := r.serverApiClient.GetPools(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	r.remotePools = make([]RemotePool, 0, len(remotePoolsResp.PoolList))

	for _, pbPool := range remotePoolsResp.PoolList {
		r.remotePools = append(r.remotePools, RemotePool{
			Name: pbPool.PoolName,
			Mode: enum.PoolMode(pbPool.PoolMode),
		})
	}

	return nil
}

type RemoteVolume struct {
	Name string
	Size uint64
	Pool string
}

type RemotePool struct {
	Name string
	Mode enum.PoolMode
}

type RemoteServerController struct {
	remoteServers []RemoteServer
}

func InitRemoteServersController(ctx context.Context) error {
	if RemoteController != nil {
		return nil
	}

	remotes, err := loadRemoteServers(ctx)
	if err != nil {
		return err
	}

	RemoteController = &RemoteServerController{
		remoteServers: remotes,
	}

	return nil
}

func (r RemoteServerController) GetRemoteServer(serverName string) (*RemoteServer, error) {
	for _, server := range r.remoteServers {
		if server.ServerName == serverName {
			return &server, nil
		}
	}

	return nil, ErrRemoteServerNotFound
}

func loadRemoteServers(ctx context.Context) ([]RemoteServer, error) {
	fileContent, err := os.ReadFile(remoteServersConfigLocation)
	if err != nil {
		return nil, err
	}

	var remotes []RemoteServer
	if err := json.Unmarshal(fileContent, &remotes); err != nil {
		return nil, err
	}

	var res []RemoteServer
	for _, server := range res {
		if err := server.InitializeRemoteServer(ctx); err != nil {
			log.Printf("Remote server %v has not beed loaded: %v", server.ConnectionString, err)
			continue
		}

		res = append(res, server)
	}

	return res, nil
}
