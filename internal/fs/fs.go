package fs

import (
	"github.com/k0tletka/SDFS/internal/fs/enum"
	"github.com/k0tletka/SDFS/internal/fs/volstorage"
)

type VolumeDispatcher struct {
	volumes []Volume
}

func NewVolumeDispatcher() (*VolumeDispatcher, error) {
	volumeSettings, err := volstorage.ExtractAllVolumeConfigs()
	if err != nil {
		return nil, err
	}
	volumes := make([]Volume, 0, len(volumeSettings))

	for _, volumeSetting := range volumeSettings {
		newVolume := Volume{}
		newVolume.ApplySettings(volumeSetting)

		volumes = append(volumes, newVolume)
	}

	return &VolumeDispatcher{
		volumes: volumes,
	}, nil
}

func (v *VolumeDispatcher) CreateNewVolume(volumeName string) (Volume, error) {
	return Volume{}, nil
}

func (v *VolumeDispatcher) GetVolume(volumeName string) (Volume, error) {
	return Volume{}, nil
}

func (v *VolumeDispatcher) UpdateVolumeInfo(volumeName string) error {
	return nil
}

func (v *VolumeDispatcher) DeleteVolume(volumeName string) error {
	return nil
}

type Volume struct {
	name        string
	storagePath string
	volumeSize  uint64
	workMode    enum.VolumeMode
}

func (v *Volume) ApplySettings(settings *volstorage.VolumeConfig) {
	v.name = settings.Name
	v.storagePath = settings.StoragePath
	v.volumeSize = settings.Size
	v.workMode = settings.Mode
}

func (v *Volume) DumpSettings() *volstorage.VolumeConfig {
	return &volstorage.VolumeConfig{
		Name:        v.name,
		StoragePath: v.storagePath,
		Size:        v.volumeSize,
		Mode:        v.workMode,
	}
}
