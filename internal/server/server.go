package server

import (
	"context"

	"github.com/k0tletka/SDFS/internal/server/clientapi"
	"github.com/k0tletka/SDFS/internal/server/config"
	"github.com/k0tletka/SDFS/internal/server/controlapi"
	"github.com/k0tletka/SDFS/internal/server/serverapi"

	"golang.org/x/sync/errgroup"
)

var (
	servers = []server{
		&clientapi.ClientAPI{},
		&controlapi.ControlAPI{},
		&serverapi.ServerAPI{},
	}
)

// server interface defined contract of all servers that runs into sdfsd
type server interface {
	Init(context.Context, *config.ServerConfig)
	StartServer() error
}

func InitAndServeServers(ctx context.Context, conf *config.ServerConfig) error {
	servGroup, ctx := errgroup.WithContext(ctx)

	for _, server := range servers {
		server.Init(ctx, conf)
		servGroup.Go(server.StartServer)
	}

	return servGroup.Wait()
}
