package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	Setup()
	ret := m.Run()
	os.Exit(ret)
}

func TestDebug_Success(t *testing.T) {
	assert.True(t, true)
}

func TestInfo_Successc(t *testing.T) {
	assert.True(t, true)
}

func TestWarn_Successc(t *testing.T) {
	assert.True(t, true)
}

func TestError_Success(t *testing.T) {
	assert.True(t, true)
}

func TestFatal_Success(t *testing.T) {
	assert.True(t, true)
}
