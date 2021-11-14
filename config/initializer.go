package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

func NewConfig() Config {
	var cfg Config
	gotenv.Load(".env")
	err := envdecode.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func Serve() {
	cfg := NewConfig()

	db, err := NewMysql(&cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
