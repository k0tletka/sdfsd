package internal

type Volume struct {
	Name        string
	StoragePath string
	VolumeSize  uint64
	Pool        string
}

func NewVolumeFromConfig(config *VolumeConfig) (*Volume, error) {
	v := &Volume{}
	v.applySettings(config)

	if err := v.checkVolumeHealth(); err != nil {
		return nil, err
	}

	return v, nil
}

func (v *Volume) ConnectVolumeToPool(pool *Pool) error {
	// TODO: Make volume connection to Pool
	v.Pool = pool.Name
	return v.onUpdate()
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

func (v *Volume) onUpdate() error {
	return SaveVolumeConfig(v.dumpSettings())
}

func (v *Volume) checkVolumeHealth() error {
	// TODO: Make checking volume health parameters on sdfsd starting
	return nil
}
