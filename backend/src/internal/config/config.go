package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port    string `yaml:"port"`
	Version string `env:"SERVERLESS_BLOG_BACKEND_VERSION"`
	YDB     YDBConfig
}

type YDBConfig struct {
	Endpoint    string `env:"SERVERLESS_BLOG_YDB_ENDPOINT"`
	AccessToken string `env:"SERVERLESS_BLOG_YDB_ACCESS_TOKEN"`
}

func Load() (*Config, error) {
	path := flag.String("config_path", "", "path to configuration file")
	flag.Parse()

	var config Config

	if err := cleanenv.ReadConfig(*path, &config); err != nil {
		return nil, fmt.Errorf("config parse error: %v", err)
	}

	return &config, nil
}
