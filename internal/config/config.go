package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

type Config struct {
	Words []string `toml:"words"`
}

func MustNew() *Config {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	c := Config{}

	_, err = toml.DecodeFile(path.Join(wd, "config", "app.toml"), &c)
	if err != nil {
		panic(err)
	}

	return &c
}
