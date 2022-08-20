package fs

import (
	"github.com/k0tletka/SDFS/internal/fs/enum"
	"github.com/k0tletka/SDFS/internal/fs/volstorage"
)

type Volume struct {
	name        string
	storagePath string
	volumeSize  uint64
	workMode    enum.VolumeMode
}

func (v *Volume) ApplySettings(settings *volstorage.VolumeConfig) {
	if settings.Name != "" {
		v.name = settings.Name
	}

	if settings.StoragePath != "" {
		v.storagePath = settings.StoragePath
	}

	if settings.Size != 0 {
		v.volumeSize = settings.Size
	}

	if settings.Mode != 0 {
		v.workMode = settings.Mode
	}
}

func (v *Volume) DumpSettings() *volstorage.VolumeConfig {
	return &volstorage.VolumeConfig{
		Name:        v.name,
		StoragePath: v.storagePath,
		Size:        v.volumeSize,
		Mode:        v.workMode,
	}
}

func (v *Volume) checkVolumeHealth() error {
	return nil
}
