package config

import (
	"errors"

	"gopkg.in/ini.v1"
)

var cfg Config

func InitConfig() (*Config, error) {
	// load config
	path := "internal/files/etc/conf/config.ini"
	cfgFile, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	if cfgFile == nil {
		return nil, errors.New("config not found")
	}

	err = cfgFile.MapTo(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetConfig() *Config {
	return &cfg
}
