package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func GetConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, err
}
