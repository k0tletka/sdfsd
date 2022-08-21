package fs

import (
	"github.com/k0tletka/SDFS/internal/fs/enum"
	"github.com/k0tletka/SDFS/internal/fs/storage"
)

type Pool struct {
	name          string
	workMode      enum.PoolMode
	syncedServers []string
}

func (p *Pool) ApplySettings(settings *storage.PoolConfig) {
	if settings.Name != "" {
		p.name = settings.Name
	}

	if settings.Mode != 0 {
		p.workMode = settings.Mode
	}

	if len(settings.SyncedServers) != 0 {
		p.syncedServers = make([]string, len(settings.SyncedServers))
		copy(p.syncedServers, settings.SyncedServers)
	}
}

func (p *Pool) DumpSettings() *storage.PoolConfig {
	syncedServers := make([]string, len(p.syncedServers))
	copy(syncedServers, p.syncedServers)

	return &storage.PoolConfig{
		Name:          p.name,
		Mode:          p.workMode,
		SyncedServers: p.syncedServers,
	}
}

func (p *Pool) syncPoolWithRemoteServers() error {
	return nil
}

func (p *Pool) checkPoolHealth() error {
	return nil
}
