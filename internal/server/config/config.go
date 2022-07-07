package config

import (
	"github.com/k0tletka/SDFS/internal/config"
	"github.com/k0tletka/SDFS/internal/fs"
)

// ServerConfig type contains all dependencies for all servers
type ServerConfig struct {
	Config        *config.RootConfig
	VolDispatcher *fs.VolumeDispatcher
}
