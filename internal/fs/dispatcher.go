package fs

import (
	"github.com/k0tletka/SDFS/internal/fs/enum"
	"github.com/k0tletka/SDFS/internal/fs/volstorage"
	"golang.org/x/sys/unix"
	"os"
)

type VolumeCreateRequest struct {
	Name         string
	StoragePath  string
	AllocateSize uint64
	WorkMode     enum.VolumeMode
}

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
		if err := newVolume.checkVolumeHealth(); err != nil {
			return nil, err
		}

		volumes = append(volumes, newVolume)
	}

	return &VolumeDispatcher{
		volumes: volumes,
	}, nil
}

func (v *VolumeDispatcher) CreateNewVolume(req VolumeCreateRequest) (Volume, error) {
	// First, check create conditions so all is okay to create new volume
	if err := v.checkCreateConditions(req); err != nil {
		return Volume{}, err
	}

	volumeConfig := &volstorage.VolumeConfig{
		Name:        req.Name,
		StoragePath: req.StoragePath,
		Size:        req.AllocateSize,
		Mode:        req.WorkMode,
	}

	newVolume := Volume{}
	newVolume.ApplySettings(volumeConfig)

	if err := volstorage.SaveVolumeConfig(volumeConfig); err != nil {
		return Volume{}, err
	}

	return newVolume, nil
}

func (v *VolumeDispatcher) GetVolume(volumeName string) (Volume, error) {
	for _, volume := range v.volumes {
		if volume.name == volumeName {
			return volume, nil
		}
	}

	return Volume{}, ErrVolumeNotFound
}

func (v *VolumeDispatcher) UpdateVolumeInfo(volumeName string) error {
	return nil
}

func (v *VolumeDispatcher) DeleteVolume(volumeName string) error {
	return nil
}

func (v *VolumeDispatcher) checkCreateConditions(req VolumeCreateRequest) error {
	if _, err := v.GetVolume(req.Name); err == nil {
		return ErrVolumeAlreadyExist
	}

	if err := v.checkCreateConditions(req); err != nil {
		return err
	}

	return nil
}

func (v *VolumeDispatcher) checkStoragePathSuitability(req VolumeCreateRequest) error {
	pathInfo, err := os.Stat(req.StoragePath)
	if err != nil {
		return err
	}

	if !pathInfo.IsDir() {
		return ErrStoragePathNotDir
	}

	files, err := os.ReadDir(req.StoragePath)
	if err != nil {
		return err
	}

	if len(files) > 0 {
		return ErrStoragePathFilesExist
	}

	var stat unix.Statfs_t
	if err := unix.Statfs(req.StoragePath, &stat); err != nil {
		return err
	}

	if stat.Bavail*uint64(stat.Bsize) < req.AllocateSize {
		return ErrStoragePathNotEnougtSize
	}

	return nil
}