package serverapi

import (
	"context"
	"github.com/k0tletka/SDFS/internal/server/config"
)

type ServerAPI struct{}

func (s *ServerAPI) Init(ctx context.Context, conf *config.ServerConfig) {}

func (s *ServerAPI) StartServer() error {
	return nil
}
