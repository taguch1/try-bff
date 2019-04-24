package middleware

import (
	"encoding/json"
	"io/ioutil"
)

// Config middleware config
type Config struct {
	Cors *CorsConfig `json:"cors"`
}

// CorsConfig cors config
type CorsConfig struct {
	AllowedOrigins []string `json:"allowedOrigins"`
	AllowedMethods []string `json:"allowedMethods"`
	AllowedHeaders []string `json:"allowedHeaders"`
}

// ConfFileName middlewares config file path
const ConfFileName = "config/middleware.json"

// NewConf load config
func NewConf(filePath string) (*Config, error) {

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	if err := json.Unmarshal(buf, config); err != nil {
		return nil, err
	}
	return config, nil
}
