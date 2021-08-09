package cmd

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigExt(t *testing.T) {
	ext := getConfigExt()
	assert.NotEmpty(t, ext)
	assert.Equal(t, ext, "json")
}

func TestGetConfigName(t *testing.T) {
	name := getConfigName()
	assert.NotEmpty(t, name)
}

func TestGetConfigDir(t *testing.T) {
	configDir, err := getConfigDir()
	assert.NoError(t, err)
	assert.NotEmpty(t, configDir)
}

func TestGetConfigPath(t *testing.T) {
	configPath, err := getConfigPath()
	assert.NoError(t, err)
	assert.NotEmpty(t, configPath)
	assert.Equal(t, ".json", filepath.Ext(configPath))
}
