package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Database struct {
	URL     string `yaml:"url"`
	Name    string `yaml:"name"`
	Timeout int    `yaml:"timeout"`
}

type Server struct {
	Debug bool   `yaml:"debug"`
	Port  string `yaml:"port"`
	Host  string `yaml:"host"`
}

type Config struct {
	DB     Database `yaml:"database"`
	Server Server   `yaml:"server"`
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
