package config

import (
	"github.com/spf13/viper"
)

type db struct {
	vpr *viper.Viper
}

func newDB(vpr *viper.Viper) *db {
	return &db{vpr: vpr}
}

func (d *db) User() string {
	return d.vpr.GetString("DB.USER")
}

func (d *db) Password() string {
	return d.vpr.GetString("DB.PASS")
}

func (d *db) Host() string {
	return d.vpr.GetString("DB.HOST")
}

func (d *db) DB() string {
	return d.vpr.GetString("DB.DATABASE")
}
