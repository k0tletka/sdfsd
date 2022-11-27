package remote

type RemoteError string

func (r RemoteError) Error() string {
	return string(r)
}

const (
	ErrRemoteServerNotInitialized RemoteError = "remote server not initialized"
	ErrRemoteAlreadyInitialized   RemoteError = "remote already initialized"
	ErrRemoteServerNotFound       RemoteError = "remote server with such name not found"
	ErrRemoteAlreadyExist         RemoteError = "remote server already exist"
	ErrServerNameCantBeEmpty      RemoteError = "remote server name cant be empty"
)
