package main

import (
	"context"
	"github.com/k0tletka/SDFS/internal/config"
	"github.com/k0tletka/SDFS/internal/server"
	srvconfig "github.com/k0tletka/SDFS/internal/server/config"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	srvConfig, err := initDependencies()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, _ := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGQUIT,
	)

	if err := server.InitAndServeServers(ctx, srvConfig); err != nil && err != context.Canceled {
		log.Fatalln(err)
	}
}

func initDependencies() (*srvconfig.ServerConfig, error) {
	conf, err := config.InitConfiguration()
	if err != nil {
		return nil, err
	}

	return &srvconfig.ServerConfig{
		Config: conf,
	}, nil
}
