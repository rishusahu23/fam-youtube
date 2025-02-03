package config

import (
	"fmt"
	"github.com/pkg/errors"
	pkgConfig "github.com/rishusahu23/fam-youtube/pkg/config"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	once   sync.Once
	config *Config
	err    error

	_, b, _, _ = runtime.Caller(0)
)

// using koanf to load config
func Load() (*Config, error) {
	once.Do(func() {
		config, err = loadConfig()
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}
	return config, err
}

func TestConfigDir() string {
	configPath := filepath.Join(b, "..")
	return configPath
}

func loadConfig() (*Config, error) {
	conf := &Config{}
	testConfigDir := TestConfigDir()
	// loads config from file
	k, _, err := pkgConfig.LoadConfig(testConfigDir, "test")
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	err = k.UnmarshalWithConf("", conf, pkgConfig.DefaultUnmarshallingConfig(conf))
	if err != nil {
		return nil, fmt.Errorf("failed to refresh config: %w", err)
	}
	return conf, nil
}

type Config struct {
	Server         *Server
	PostgresConfig *PostgresConfig
	ApiKeys        []string
	Environment    string
}

type Server struct {
	Port         int
	GrpcPort     int
	GrpcHttpPort int
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}
