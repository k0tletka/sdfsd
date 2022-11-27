package fs

import (
	"github.com/k0tletka/sdfsd/internal/fs/enum"
	fsinternal "github.com/k0tletka/sdfsd/internal/fs/internal"
	"golang.org/x/sys/unix"
	"os"
)

type VolumeCreateRequest struct {
	Name         string
	StoragePath  string
	AllocateSize uint64
	PoolName     string
}

type PoolCreateRequest struct {
	PoolName string
	Mode     enum.PoolMode
}

type VolumeDispatcher struct {
	volumes []fsinternal.Volume
	pools   []fsinternal.Pool
}

func NewVolumeDispatcher() (*VolumeDispatcher, error) {
	volumeSettings, err := fsinternal.ExtractAllVolumeConfigs()
	if err != nil {
		return nil, err
	}
	volumes := make([]fsinternal.Volume, 0, len(volumeSettings))

	for _, volumeSetting := range volumeSettings {
		newVolume, err := fsinternal.NewVolumeFromConfig(volumeSetting)
		if err != nil {
			return nil, err
		}
		volumes = append(volumes, *newVolume)
	}

	poolSettings, err := fsinternal.ExtractAllPoolConfigs()
	if err != nil {
		return nil, err
	}
	pools := make([]fsinternal.Pool, 0, len(poolSettings))

	for _, poolSetting := range poolSettings {
		newPool, err := fsinternal.NewPoolFromConfig(poolSetting)
		if err != nil {
			return nil, err
		}
		pools = append(pools, *newPool)
	}

	return &VolumeDispatcher{
		volumes: volumes,
		pools:   pools,
	}, nil
}

func (v *VolumeDispatcher) CreateNewVolume(req VolumeCreateRequest) (fsinternal.Volume, error) {
	// First, check create conditions so all is okay to create new volume
	if err := v.checkCreateConditions(req); err != nil {
		return fsinternal.Volume{}, err
	}

	newVolume := fsinternal.Volume{
		Name:        req.Name,
		StoragePath: req.StoragePath,
		VolumeSize:  req.AllocateSize,
	}

	if req.PoolName != "" {
		pool, err := v.GetPool(req.PoolName)
		if err != nil {
			return fsinternal.Volume{}, err
		}

		if err := newVolume.ConnectVolumeToPool(&pool); err != nil {
			return fsinternal.Volume{}, err
		}
	}

	v.volumes = append(v.volumes, newVolume)
	return newVolume, nil
}

func (v *VolumeDispatcher) CreateNewPool(req PoolCreateRequest) (fsinternal.Pool, error) {
	if _, err := v.GetPool(req.PoolName); err == nil {
		return fsinternal.Pool{}, ErrPoolAlreadyExist
	}

	newPool := fsinternal.Pool{
		Name:     req.PoolName,
		WorkMode: req.Mode,
	}

	if err := newPool.SyncPoolWithRemoteServers(); err != nil {
		return fsinternal.Pool{}, err
	}

	v.pools = append(v.pools, newPool)
	return newPool, nil
}

func (v *VolumeDispatcher) GetVolume(volumeName string) (fsinternal.Volume, error) {
	for _, volume := range v.volumes {
		if volume.Name == volumeName {
			return volume, nil
		}
	}

	return fsinternal.Volume{}, ErrNotFound
}

func (v *VolumeDispatcher) GetVolumes() []fsinternal.Volume {
	res := make([]fsinternal.Volume, len(v.volumes))
	copy(res, v.volumes)
	return res
}

func (v *VolumeDispatcher) GetPool(poolName string) (fsinternal.Pool, error) {
	for _, pool := range v.pools {
		if pool.Name == poolName {
			return pool, nil
		}
	}

	return fsinternal.Pool{}, ErrNotFound
}

func (v *VolumeDispatcher) GetPools() []fsinternal.Pool {
	res := make([]fsinternal.Pool, len(v.pools))
	copy(res, v.pools)
	return res
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
