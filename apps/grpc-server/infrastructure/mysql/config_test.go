package mysql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConf_Success(t *testing.T) {
	assert.True(t, true)
	config, err := NewConf("../..//config/mysql.json")
	assert.Nil(t, err)
	assert.Equal(t, "tcp(127.0.0.1:3306)", config.ConnectionName)
	assert.Equal(t, "root", config.User)
	assert.Equal(t, "rdbms-password", config.Password)
}
