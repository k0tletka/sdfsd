package fs

import (
	"github.com/k0tletka/sdfsd/internal/fs/enum"
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
	volumes []Volume
	pools   []Pool
}

func NewVolumeDispatcher() (*VolumeDispatcher, error) {
	volumeSettings, err := extractAllVolumeConfigs()
	if err != nil {
		return nil, err
	}
	volumes := make([]Volume, 0, len(volumeSettings))

	for _, volumeSetting := range volumeSettings {
		newVolume := Volume{}
		newVolume.applySettings(volumeSetting)
		if err := newVolume.checkVolumeHealth(); err != nil {
			return nil, err
		}

		volumes = append(volumes, newVolume)
	}

	poolSettings, err := extractAllPoolConfigs()
	if err != nil {
		return nil, err
	}
	pools := make([]Pool, 0, len(poolSettings))

	for _, poolSetting := range poolSettings {
		newPool := Pool{}
		newPool.applySettings(poolSetting)
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

	newVolume := Volume{
		Name:        req.Name,
		StoragePath: req.StoragePath,
		VolumeSize:  req.AllocateSize,
	}

	if req.PoolName != "" {
		pool, err := v.GetPool(req.PoolName)
		if err != nil {
			return Volume{}, err
		}

		if err := newVolume.connectVolumeToPool(&pool); err != nil {
			return Volume{}, err
		}
	}

	if err := saveVolumeConfig(newVolume.dumpSettings()); err != nil {
		return Volume{}, err
	}

	return newVolume, nil
}

func (v *VolumeDispatcher) CreateNewPool(req PoolCreateRequest) (Pool, error) {
	if _, err := v.GetPool(req.PoolName); err == nil {
		return Pool{}, ErrPoolAlreadyExist
	}

	newPool := Pool{
		Name:     req.PoolName,
		WorkMode: req.Mode,
	}

	if err := newPool.syncPoolWithRemoteServers(); err != nil {
		return Pool{}, err
	}
	if err := savePoolConfig(newPool.dumpSettings()); err != nil {
		return Pool{}, err
	}

	return newPool, nil
}

func (v *VolumeDispatcher) GetVolume(volumeName string) (Volume, error) {
	for _, volume := range v.volumes {
		if volume.Name == volumeName {
			return volume, nil
		}
	}

	return Volume{}, ErrNotFound
}

func (v *VolumeDispatcher) GetVolumes() []Volume {
	res := make([]Volume, len(v.volumes))
	copy(res, v.volumes)
	return res
}

func (v *VolumeDispatcher) GetPool(poolName string) (Pool, error) {
	for _, pool := range v.pools {
		if pool.Name == poolName {
			return pool, nil
		}
	}

	return Pool{}, ErrNotFound
}

func (v *VolumeDispatcher) GetPools() []Pool {
	res := make([]Pool, len(v.pools))
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
