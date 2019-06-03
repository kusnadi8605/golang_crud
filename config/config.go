package config

import (
	"os"
	"promo_api/parser"
)

//Param ..
var Param Configuration

// Configuration stores global configuration loaded from json file
type Configuration struct {
	ListenPort string            `yaml:"listenPort"`
	Query      string            `yaml:"query"`
	DBUrl      string            `yaml:"dbUrl"`
	DBType     string            `yaml:"dbType"`
	RedisURL   string            `yaml:"redisURL"`
	LogDir     string            `yaml:"logDir"`
	LogsFile   map[string]string `yaml:"logsFile"`
	ListIP     []string          `yaml:"listIP"`
}

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := parser.LoadYAML(fn, &Param); err != nil {
		os.Exit(1)
	}
}
