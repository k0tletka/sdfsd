package clientapi

import (
	"context"
	"github.com/k0tletka/SDFS/internal/server/config"
)

type ClientAPI struct{}

func (s *ClientAPI) Init(ctx context.Context, config *config.ServerConfig) {}

func (s *ClientAPI) StartServer() error {
	return nil
}
