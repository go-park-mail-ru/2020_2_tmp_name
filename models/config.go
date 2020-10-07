package models

import (
	"time"

	"github.com/BurntSushi/toml"
)

var (
	configPath = "config/config.toml"
)

type duration time.Duration

type Config struct {
	SQLDataBase SQLDataBase `toml:"SQLDataBase"`
}

func LoadConfig(c *Config) {
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		return
	}
}
