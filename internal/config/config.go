package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	NatsStreaming `yaml:"nats_streaming"`
	HTTPServer    `yaml:"http_server"`
	DB            `yaml:"database"`
}

type NatsStreaming struct {
	ClusterID   string `yaml:"cluster_id"`
	ClientID    string `yaml:"client_id"`
	ChannelName string `yaml:"channel_name"`
	QueueGroup  string `yaml:"queue_group"`
	URL         string `yaml:"url"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	TimeoutIdle time.Duration `yaml:"timeout_idle"`
}
type DB struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Database string `yaml:"database"`
}

func Load() (*Config, error) {
	log.Println("Loading config...")
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
