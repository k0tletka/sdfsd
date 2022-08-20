package fs

import (
	"github.com/k0tletka/SDFS/internal/fs/storage"
	"golang.org/x/sys/unix"
	"os"
)

type VolumeCreateRequest struct {
	Name         string
	StoragePath  string
	AllocateSize uint64
	PoolName     string
}

type VolumeDispatcher struct {
	volumes []Volume
	pools   []Pool
}

func NewVolumeDispatcher() (*VolumeDispatcher, error) {
	volumeSettings, err := storage.ExtractAllVolumeConfigs()
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

	poolSettings, err := storage.ExtractAllPoolConfigs()
	if err != nil {
		return nil, err
	}
	pools := make([]Pool, 0, len(poolSettings))

	for _, poolSetting := range poolSettings {
		newPool := Pool{}
		newPool.ApplySettings(poolSetting)
		if err := newPool.checkPoolHealth(); err != nil {
			return nil, err
		}

		pools = append(pools, newPool)
	}

	return &VolumeDispatcher{
		volumes: volumes,
		pools:   pools,
	}, nil
}

func (v *VolumeDispatcher) CreateNewVolume(req VolumeCreateRequest) (Volume, error) {
	// First, check create conditions so all is okay to create new volume
	if err := v.checkCreateConditions(req); err != nil {
		return Volume{}, err
	}

	volumeConfig := &storage.VolumeConfig{
		Name:        req.Name,
		StoragePath: req.StoragePath,
		Size:        req.AllocateSize,
	}

	newVolume := Volume{}
	newVolume.ApplySettings(volumeConfig)

	if req.PoolName != "" {
		pool, err := v.GetPool(req.PoolName)
		if err != nil {
			return Volume{}, err
		}

		if err := newVolume.ConnectVolumeToPool(&pool); err != nil {
			return Volume{}, err
		}
	}

	if err := storage.SaveVolumeConfig(volumeConfig); err != nil {
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

	return Volume{}, ErrNotFound
}

func (v *VolumeDispatcher) GetPool(poolName string) (Pool, error) {
	for _, pool := range v.pools {
		if pool.name == poolName {
			return pool, nil
		}
	}

	return Pool{}, ErrNotFound
}

func (v *VolumeDispatcher) checkCreateConditions(req VolumeCreateRequest) error {
	if _, err := v.GetVolume(req.Name); err == nil {
		return ErrVolumeAlreadyExist
	}

	if err := v.checkStoragePathSuitability(req); err != nil {
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
