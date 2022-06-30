package config

// RootConfig is a main struct that contains all config values from toml file
type RootConfig struct {
	ClientAPIConf  ClientAPIConfig  `toml:"clientapi"`
	ControlAPIConf ControlAPIConfig `toml:"controlapi"`
	ServerAPIConf  ServerAPIConfig  `toml:"serverapi"`
}

type ClientAPIConfig struct {
}

type ControlAPIConfig struct {
	ListenAddr   string `toml:"listenaddr"`
	ListenPort   uint16 `toml:"listenport"`
	UseSSL       bool   `toml:"use_ssl"`
	CertFilePath string `toml:"cert_filepath"`
	KeyFilePath  string `toml:"key_filepath"`
}

type ServerAPIConfig struct {
}

func InitConfiguration() (*RootConfig, error) {
	return defaultConfiguration(), nil
}

func defaultConfiguration() *RootConfig {
	return &RootConfig{
		ClientAPIConf: ClientAPIConfig{},
		ControlAPIConf: ControlAPIConfig{
			ListenAddr: "localhost",
			ListenPort: 16500,
			UseSSL:     false,
		},
		ServerAPIConf: ServerAPIConfig{},
	}
}
