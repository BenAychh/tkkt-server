package config

import "github.com/spf13/viper"

type http struct {
	vpr *viper.Viper
}

func newHTTP(vpr *viper.Viper) *http {
	return &http{vpr: vpr}
}

func (h *http) Port() string {
	return h.vpr.GetString("HTTP.PORT")
}
