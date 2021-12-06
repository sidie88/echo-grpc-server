package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type OmDbApiConfig struct {
	ApiKey string `envconfig:"OMDB_APIKEY" default:"faf7e5bb"`
}

func NewOmDbApiConfig() *OmDbApiConfig {
	spec := &OmDbApiConfig{}
	err := envconfig.Process("OMDB", spec)
	if err != nil {
		log.Error(`Invalid "OMDB"" config`)
	}
	return spec
}