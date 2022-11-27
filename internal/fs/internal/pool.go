package internal

import (
	"github.com/k0tletka/sdfsd/internal/fs/enum"
)

type Pool struct {
	Name     string
	WorkMode enum.PoolMode

	syncedServers []string
}

func NewPoolFromConfig(config *PoolConfig) (*Pool, error) {
	p := &Pool{}
	p.applySettings(config)

	// Check that all alright with pool
	if err := p.checkPoolHealth(); err != nil {
		return nil, err
	}

	// Sync pool with other servers
	if err := p.SyncPoolWithRemoteServers(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Pool) GetSyncedServers() []string {
	res := make([]string, len(p.syncedServers))
	copy(res, p.syncedServers)
	return res
}

func (p *Pool) SyncPoolWithRemoteServers() error {
	// TODO: Make Pool syncing with other servers on creation
	return p.onUpdate()
}

func (p *Pool) applySettings(settings *PoolConfig) {
	if settings.Name != "" {
		p.Name = settings.Name
	}

	if settings.Mode != 0 {
		p.WorkMode = settings.Mode
	}

	if len(settings.SyncedServers) != 0 {
		p.syncedServers = make([]string, len(settings.SyncedServers))
		copy(p.syncedServers, settings.SyncedServers)
	}
}

func (p *Pool) dumpSettings() *PoolConfig {
	syncedServers := make([]string, len(p.syncedServers))
	copy(syncedServers, p.syncedServers)

	return &PoolConfig{
		Name:          p.Name,
		Mode:          p.WorkMode,
		SyncedServers: p.syncedServers,
	}
}

func (p *Pool) onUpdate() error {
	return SavePoolConfig(p.dumpSettings())
}

func (p *Pool) checkPoolHealth() error {
	// TODO: Make check Pool health on sdfsd starting
	return nil
}
