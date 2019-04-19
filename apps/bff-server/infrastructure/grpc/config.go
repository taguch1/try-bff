package grpc

import (
	"encoding/json"
	"io/ioutil"
)

// Config grpc config
type Config struct {
	TargetAddress string `json:"targetAddress"`
	TimeoutMillis int    `json:"timeoutMillis"`
}

// ConfFileName grpc config file path
const ConfFileName = "config/grpc.json"

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
