package fs

type FSError string

func (f FSError) Error() string {
	return string(f)
}

const (
	ErrNotFound FSError = "requested object has not beed found"

	ErrVolumeAlreadyExist FSError = "volume with such name already exist"

	ErrPoolAlreadyExist FSError = "pool with suck name already exist"

	ErrStoragePathNotDir        FSError = "passed path is not a dir"
	ErrStoragePathFilesExist    FSError = "in storage path must not be any files"
	ErrStoragePathNotEnougtSize FSError = "there is not enought size for new volume"
)
