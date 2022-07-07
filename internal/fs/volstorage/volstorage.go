package volstorage

import (
	"encoding/gob"
	"github.com/k0tletka/SDFS/internal/fs/enum"
	"os"
)

const (
	volumeObjBasePath = "/var/sdfsd/"
)

type VolumeConfig struct {
	Name        string
	StoragePath string
	Size        uint64
	Mode        enum.VolumeMode
}

func ExtractAllVolumeConfigs() ([]*VolumeConfig, error) {
	dir, err := os.ReadDir(volumeObjBasePath)
	if err != nil {
		return nil, err
	}

	res := make([]*VolumeConfig, 0, len(dir))

	for _, entry := range dir {
		volConf, err := extractVolumeConfigByPath(volumeObjBasePath + entry.Name())
		if err != nil {
			return nil, err
		}

		res = append(res, volConf)
	}

	return res, nil
}

func ExtractVolumeConfig(volumeName string) (*VolumeConfig, error) {
	return extractVolumeConfigByPath(volumeObjBasePath + volumeName + ".gob")
}

func SaveVolumeConfig(volumeConfig *VolumeConfig) error {
	objFile, err := os.OpenFile(
		volumeObjBasePath+volumeConfig.Name+".gob",
		os.O_WRONLY|os.O_CREATE|os.O_SYNC,
		0600,
	)

	if err != nil {
		return err
	}
	defer objFile.Close()

	volConfigEncoder := gob.NewEncoder(objFile)
	return volConfigEncoder.Encode(volumeConfig)
}

func extractVolumeConfigByPath(volumeConfigPath string) (*VolumeConfig, error) {
	objFile, err := os.OpenFile(volumeConfigPath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer objFile.Close()

	volumeConfig := VolumeConfig{}

	volConfigDecoder := gob.NewDecoder(objFile)
	if err := volConfigDecoder.Decode(&volumeConfig); err != nil {
		return nil, err
	}

	return &volumeConfig, nil
}
