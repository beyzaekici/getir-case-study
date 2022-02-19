package data

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	MongoServer   string
	MongoDatabase string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
