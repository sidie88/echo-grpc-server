package config

import "time"

type MySqlConfig struct {
	UserName        string        `envconfig:"SQL_USERNAME" default:"root"`
	Password        string        `envconfig:"SQL_PASSWORD"`
	ConnMaxLifeTime time.Duration `envconfig:"SQL_CONN_MAX_LIFE" default:"120s"`
	MaxOpenConns    int32         `envconfig:"SQL_MAX_OPEN_CONNS" default:"10"`
	MaxIdleConns    int32         `envconfig:"SQL_MAX_IDLE_CONNS" default:"10"`
}
