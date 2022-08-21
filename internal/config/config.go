package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	defaultConfigFileLocation = "config.toml"
)

// RootConfig is a main struct that contains all config values from toml file
type RootConfig struct {
	ServerName     string           `toml:"server_name"`
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
	ListenAddr string `toml:"listenaddr"`
	ListenPort uint16 `toml:"listenport"`
}

func InitConfiguration() *RootConfig {
	conf := defaultConfiguration()

	confLocation, exist := os.LookupEnv("CONFFILE")
	if !exist {
		confLocation = defaultConfigFileLocation
	}

	if _, err := toml.DecodeFile(confLocation, &conf); err != nil {
		log.Println("WARN: Can't read config file: ", err)
		log.Println("Using default config values")
	}

	return conf
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
