package remote

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"
)

var (
	RemoteController *RemoteServerController
)

const (
	remoteServersConfigLocation = "/var/sdfds/remoteServers.json"
)

type RemoteServer struct {
	ConnectionString string
	ServerName       string

	conn          net.Conn // TODO: use gRPC
	isInitialized bool
	remoteVolumes []RemoteVolume
}

func (r *RemoteServer) IsInitialized() bool {
	return r.isInitialized
}

func (r *RemoteServer) InitializeRemoteServer(ctx context.Context) error {
	conn, err := net.Dial("tcp", r.ConnectionString)
	if err != nil {
		return err
	}

	r.conn = conn
	return r.LoadRemoteVolumes()
}

func (r *RemoteServer) LoadRemoteVolumes() error {
	if !r.IsInitialized() {
		return ErrRemoteServerNotInitialized
	}

	// TODO: make request to get all volumes from remote server
	return nil
}

type RemoteVolume struct {
	ConnectionString string
	Name             string
	Size             uint64
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
