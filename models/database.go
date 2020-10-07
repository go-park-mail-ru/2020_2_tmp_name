package models

import "time"

// SQLDataBase struct
type SQLDataBase struct {
	Server          string        `toml:"Server"`
	Database        string        `toml:"Database"`
	ApplicationName string        `toml:"ApplicationName"`
	MaxIdleConns    int           `toml:"MaxIdleConns"`
	MaxOpenConns    int           `toml:"MaxOpenConns"`
	ConnMaxLifetime time.Duration `toml:"ConnMaxLifetime"`
	UserID          string
	Password        string
}
