package remote

type RemoteError string

func (r RemoteError) Error() string {
	return string(r)
}

const (
	ErrRemoteServerNotInitialized RemoteError = "remote server not initialized"

	ErrRemoteServerNotFound RemoteError = "remote server with such name not found"
)
