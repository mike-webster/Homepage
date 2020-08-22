package env

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"homepage/keys"

	yaml "gopkg.in/yaml.v1"
)

var (
	configPath = "env/app.yaml"
	_cfg       = &Config{}
)

type Config struct {
	Port    string   `yaml:"port"`
	Teams   []string `yaml:"teams"`
	Leagues []string `yaml:"leagues"`
}

func GetConfig() (*Config, error) {
	return loadConfig()
}

func GetConfigFromContext(ctx context.Context) (*Config, error) {
	cfg := ctx.Value(keys.AppConfig)
	if cfg == nil {
		cfg = ctx.Value(string(keys.AppConfig))
		if cfg == nil {
			return loadConfig()
		}

		config, ok := cfg.(*Config)
		if !ok {
			return loadConfig()
		}
		return config, nil
	}

	config, ok := cfg.(*Config)
	if !ok {
		return loadConfig()
	}
	return config, nil
}

func loadConfig() (*Config, error) {
	// TODO: figure out how to cache this
	path, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	bp, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}

	bp = filepath.Dir(bp)

	var configs map[string]Config
	if err = yaml.Unmarshal(f, &configs); err != nil {
		return nil, err
	}

	ret := configs["development"]
	return &ret, nil
}
