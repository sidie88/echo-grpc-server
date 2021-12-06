package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type GRpcConfig struct {
	Address string        `envconfig:"GRPC_ADDRESS" default:"localhost"`
	Port    string        `envconfig:"GRPC_PORT" default:"8888"`
}

func NewGRpcConfig() *GRpcConfig {
	spec := &GRpcConfig{}
	err := envconfig.Process("GRPC", spec)
	if err != nil {
		log.Error(`Invalid "GRPC"" config`)
	}
	return spec
}