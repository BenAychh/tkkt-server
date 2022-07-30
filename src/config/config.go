package config

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	vpr *viper.Viper
}

func newConfig() (*Config, error) {
	vpr := viper.New()
	vpr.SetEnvPrefix("tkkt")
	vpr.AllowEmptyEnv(false)
	vpr.AutomaticEnv()
	vpr.SetDefault("HTTP.PORT", "8080")
	if vpr.GetBool("USE_CONSUL") {
		err := vpr.AddRemoteProvider("consul", vpr.GetString("CONSUL_ADDR"), vpr.GetString("CONSUL_KEY_PREFIX"))
		if err != nil {
			return nil, err
		}
		vpr.SetConfigType("json")
		err = vpr.ReadRemoteConfig()
	}
	return &Config{vpr: vpr}, nil
}

func (c *Config) DB() *db {
	return newDB(c.vpr)
}

func (c *Config) HTTP() *http {
	return newHTTP(c.vpr)
}
