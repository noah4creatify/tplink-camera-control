package pkg

import (
	"os"

	"github.com/pelletier/go-toml"
)

type ConfigOptions struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	Address  string `json:"address"`
	Stok     string `json:"stok"`
}

func (c *ConfigOptions) UpdateSaveStok(stok string) error {
	c.Stok = stok
	return WriteConfig(c)
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

func GetConfig() (*ConfigOptions, error) {
	cfg, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	if cfg.UserName == "" {
		cfg.UserName = DefaultUserName
	}
	if cfg.Address == "" {
		return nil, ErrAddressNotSet
	}
	if cfg.PassWord == "" {
		return nil, ErrPassWordNotSet
	}
	return cfg, nil
}

func WriteConfig(conf *ConfigOptions) error {
	configData, err := toml.Marshal(conf)
	if err != nil {
		return err
	}
	return os.WriteFile(CONFIGFILE, configData, 0644)
}
