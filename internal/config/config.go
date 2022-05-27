package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var PathConfigFile string

type Server struct {
	Port int `yaml:"port"`
}

type MongoDB struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (m *MongoDB) URI() string {
	return fmt.Sprintf("mongodb://%s:%s/", m.Host, strconv.Itoa(m.Port))
}

type Config struct {
	Server  `yaml:"server"`
	MongoDB `yaml:"mongo"`
}

func init() {
	flag.StringVar(&PathConfigFile, "confFile", "./configs/configs.yaml", "path to config file")
}

func New(path string) (*Config, error) {

	flag.Parse()

	cfg := &Config{}

	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(cfgFile, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
