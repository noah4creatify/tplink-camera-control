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
}

func ReadConfig() (conf *ConfigOptions, err error) {
	if err = CheckOrCreateConfigFile(); err != nil {
		return nil, ErrReadConfigFile
	}
	configData, err := os.ReadFile(CONFIGFILE)
	if err == nil {
		err = toml.Unmarshal(configData, conf)
	}
	return conf, err
}

func WriteConfig(conf *ConfigOptions) error {
	configData, err := toml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(CONFIGFILE, configData, 0644)
}
