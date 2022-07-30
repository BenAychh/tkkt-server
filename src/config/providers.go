package config

import (
	"log"
	"sync"
)

var (
	cfg     *Config
	cfgOnce sync.Once
)

func ProvideConfig() *Config {
	cfgOnce.Do(func() {
		var err error
		cfg, err = newConfig()
		if err != nil {
			log.Fatalf("failed to create new config %v", err)
		}
	})
	return cfg
}
