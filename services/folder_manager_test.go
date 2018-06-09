package services

import (
	"testing"
	"github.com/alicebob/miniredis"
	"time"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	RedisClient miniredis.RedisDB
}

func (m MockClient) Get(key string) (string, error) {
	return m.RedisClient.Get(key)

}
func (m MockClient) Set(key, value string, lifeTime time.Duration) error {
	return m.RedisClient.Set(key, value)
}
func (m MockClient) Delete(key string) error {
	deleted := m.RedisClient.Del(key)
	if !deleted {
		return errors.New("Unable to delete")
	}
	return nil
}
func TestNewFolderWatcher(t *testing.T) {
	m := MockClient{}
	w := &Watcher{}
	fw := NewFolderWatcher(m, w)
	assert.NotEmpty(t, fw)
}
func TestFolderWatcher_StartWatcherInvalidFolder(t *testing.T) {
	m := MockClient{}
	w := &Watcher{}
	fw := NewFolderWatcher(m, w)
	_, err := fw.StartWatcher("test", time.Duration(0))
	assert.Error(t, err)
}
func TestFolderWatcher_StartWatcherValidFolderInvalidDuration(t *testing.T) {
	m := MockClient{}
	w := &Watcher{}
	fw := NewFolderWatcher(m, w)
	_, err := fw.StartWatcher(".", time.Duration(0))
	assert.Error(t, err)
}
func TestFolderWatcher_StartWatcher(t *testing.T) {
	m := MockClient{}
	w := &Watcher{}
	fw := NewFolderWatcher(m, w)
	defer fw.Watcher.Close()
	watcher, err := fw.StartWatcher(".", time.Duration(1))
	assert.NoError(t, err)
	assert.NotEmpty(t, watcher)
}