package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type Config struct {
	NatsStreaming `yaml:"nats_streaming"`
	HTTPServer    `yaml:"http_server"`
}

type NatsStreaming struct {
	ClusterID   string `yaml:"cluster_id"`
	ClientID    string `yaml:"client_id"`
	ChannelName string `yaml:"channel_name"`
	URL         string `yaml:"url"`
}

type HTTPServer struct {
	Address string        `yaml:"address"`
	Timeout time.Duration `yaml:"timeout"`
}

func Load() (*Config, error) {
	if _, err := os.Stat("config.yaml"); os.IsNotExist(err) {
		return nil, err
	}
	configData, err := os.ReadFile("config.yaml")

	var cfg Config
	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
