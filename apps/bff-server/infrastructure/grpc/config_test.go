package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConf_Success(t *testing.T) {
	assert.True(t, true)
	config, err := NewConf("../../config/grpc.json")
	assert.Nil(t, err)
	assert.Equal(t, "localhost:50051", config.TargetAddress)
	assert.Equal(t, 3000, config.TimeoutMillis)
}
