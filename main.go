package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/k0tletka/SDFS/internal/config"
	"github.com/k0tletka/SDFS/internal/server"
	srvconfig "github.com/k0tletka/SDFS/internal/server/config"
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
	conf := config.InitConfiguration()

	return &srvconfig.ServerConfig{
		Config: conf,
	}, nil
}
