package mysql

import (
	"encoding/json"
	"io/ioutil"
)

// Config grpc config
type Config struct {
	ConnectionName string `json:"connectionName"`
	User           string `json:"user"`
	Password       string `json:"password"`
}

// ConfFileName grpc config file path
const ConfFileName = "config/mysql.json"

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
