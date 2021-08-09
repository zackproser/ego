package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigExt(t *testing.T) {
	ext := getConfigExt()
	assert.NotEmpty(t, ext)
	assert.Equal(t, ext, ".json")
}

func TestGetConfigName(t *testing.T) {
	name := getConfigName()
	assert.NotEmpty(t, name)
}
