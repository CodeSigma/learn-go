package config

import (
	"log"

	"github.com/go-ini/ini"
)

type Database struct {
	Username string
	Password string
	Host     string
	Port     int
}

var DatabaseConfig = &Database{}

var cfg *ini.File

func InitConfig() {
	var err error
	cfg, err = ini.Load("config/conf.ini")

	if err != nil {
		log.Fatalf("Failed to load config")
	}

	mapConfig("database", DatabaseConfig)
}

func mapConfig(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Failed to map config %s", section)
	}
}
