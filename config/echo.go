package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type EchoConfig struct {
	Port    string        `envconfig:"ECHO_PORT" default:"8080"`
}

func NewEchoConfig() *EchoConfig {
	spec := &EchoConfig{}
	err := envconfig.Process("ECHO", spec)
	if err != nil {
		log.Error(`Invalid "ECHO"" config`)
	}
	return spec
}
