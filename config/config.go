package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type (
	AppConfig struct {
		Debug    bool   `envconfig:"debug"`
		HttpPort int    `envconfig:"http_port"`
		HttpHost string `envconfig:"http_host"`

		DBConfig struct {
			DbDriver   string `envconfig:"driver"`
			DbPort     string `envconfig:"port"`
			DbHost     string `envconfig:"host"`
			DbName     string `envconfig:"name"`
			DbUser     string `envconfig:"user"`
			DbPassword string `envconfig:"password"`
			DbDebug    int    `envconfig:"debug"`

			MaxConnLifetimeSeconds int `envconfig:"max_conn_lifetime_seconds"`
			MaxOpenConns           int `envconfig:"max_open_conns"`
			MaxIdleConns           int `envconfig:"max_idle_conns"`
		} `envconfig:"db"`
	}
)

var (
	cfg *AppConfig
)

func LoadConfig() *AppConfig {
	var xfg AppConfig
	err := envconfig.Process(AppName, &xfg)
	if err != nil {
		panic(fmt.Sprintf("cannot read in config: %s", err))
	}

	cfg = &xfg
	return cfg
}
