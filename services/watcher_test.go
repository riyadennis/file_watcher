package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateWatcher(t *testing.T) {
	watcher := CreateWatcher()
	assert.NotNil(t, watcher)
}