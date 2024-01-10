package config

import (
	"encoding/base64"
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
	Endpoint                 string `env:"SERVERLESS_BLOG_YDB_ENDPOINT"`
	ServiceAccountKey        string
	ServiceAccountKeyEncoded string `env:"SERVERLESS_BLOG_YDB_SERVICE_ACCOUNT_KEY_ENCODED"`
}

func Load() (*Config, error) {
	path := flag.String("config_path", "", "path to configuration file")
	flag.Parse()

	var config Config

	if err := cleanenv.ReadConfig(*path, &config); err != nil {
		return nil, fmt.Errorf("config parse error: %v", err)
	}

	serviceAccountKey, err := base64.StdEncoding.DecodeString(config.YDB.ServiceAccountKeyEncoded)
	if err != nil {
		return nil, fmt.Errorf("service account key decode error: %w", err)
	}

	config.YDB.ServiceAccountKey = string(serviceAccountKey)

	return &config, nil
}
