package fs

type Volume struct {
	Name        string
	StoragePath string
	VolumeSize  uint64
	Pool        string
}

func (v *Volume) applySettings(settings *VolumeConfig) {
	if settings.Name != "" {
		v.Name = settings.Name
	}

	if settings.StoragePath != "" {
		v.StoragePath = settings.StoragePath
	}

	if settings.Size != 0 {
		v.VolumeSize = settings.Size
	}

	if settings.Pool != "" {
		v.Pool = settings.Pool
	}
}

func (v *Volume) dumpSettings() *VolumeConfig {
	return &VolumeConfig{
		Name:        v.Name,
		StoragePath: v.StoragePath,
		Size:        v.VolumeSize,
		Pool:        v.Pool,
	}
}

func (v *Volume) connectVolumeToPool(pool *Pool) error {
	// TODO: Make volume connection to Pool
	return nil
}

func (v *Volume) checkVolumeHealth() error {
	// TODO: Make checking volume health parameters on sdfsd starting
	return nil
}
