package pkg

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml"
)

var (
	ErrReadConfigFile = errors.New("read config file error")
)

type ConfigOptions struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Address  string `json:"address"`
	PubKey   string `json:"pub_key"`
}

func ReadConfig() (*ConfigOptions, error) {
	if err := CheckOrCreateConfigFile(); err != nil {
		return nil, ErrReadConfigFile
	}
	configData, err := os.ReadFile(CONFIGFILE)
	var cfg ConfigOptions
	if err == nil {
		err = toml.Unmarshal(configData, &cfg)
	}
	return &cfg, err
}

func WriteConfig(conf *ConfigOptions) error {
	configData, err := toml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(CONFIGFILE, configData, 0644)
}
