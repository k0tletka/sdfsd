package fs

import (
	"github.com/k0tletka/sdfsd/internal/fs/storage"
)

type Volume struct {
	name        string
	storagePath string
	volumeSize  uint64
	pool        string
}

func (v *Volume) ApplySettings(settings *storage.VolumeConfig) {
	if settings.Name != "" {
		v.name = settings.Name
	}

	if settings.StoragePath != "" {
		v.storagePath = settings.StoragePath
	}

	if settings.Size != 0 {
		v.volumeSize = settings.Size
	}

	if settings.Pool != "" {
		v.pool = settings.Pool
	}
}

func (v *Volume) DumpSettings() *storage.VolumeConfig {
	return &storage.VolumeConfig{
		Name:        v.name,
		StoragePath: v.storagePath,
		Size:        v.volumeSize,
		Pool:        v.pool,
	}
}

func (v *Volume) connectVolumeToPool(pool *Pool) error {
	// TODO: Make volume connection to pool
	return nil
}

func (v *Volume) checkVolumeHealth() error {
	// TODO: Make checking volume health parameters on sdfsd starting
	return nil
}
