package main

import (
	"context"
	"github.com/k0tletka/sdfsd/internal/fs"
	"github.com/k0tletka/sdfsd/internal/remote"
	"log"
	"os/signal"
	"syscall"

	"github.com/k0tletka/sdfsd/internal/config"
	"github.com/k0tletka/sdfsd/internal/server"
	srvconfig "github.com/k0tletka/sdfsd/internal/server/config"
)

func main() {
	ctx, _ := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)

	if err := remote.InitRemoteServersController(ctx); err != nil {
		log.Fatalln(err)
	}

	srvConfig, err := initDependencies()
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.InitAndServeServers(ctx, srvConfig); err != nil && err != context.Canceled {
		log.Fatalln(err)
	}
}

func initDependencies() (*srvconfig.ServerConfig, error) {
	conf := config.InitConfiguration()
	volDispatcher, err := fs.NewVolumeDispatcher()
	if err != nil {
		return nil, err
	}

	return &srvconfig.ServerConfig{
		Config:        conf,
		VolDispatcher: volDispatcher,
	}, nil
}
