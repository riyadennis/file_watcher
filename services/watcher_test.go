package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestWatcher_AddInvalidFolder(t *testing.T) {
	w := &Watcher{}
	err := w.Add("invalid")
	assert.Error(t, err)
}
func TestWatcher_AddValidEmptyFolder(t *testing.T) {
	w := createWatcher()
	emptyFolder := "empty"
	os.Mkdir(emptyFolder, os.ModeDir)
	defer os.Remove(emptyFolder)
	err := w.Add(emptyFolder)
	assert.Error(t, err)
}
func TestWatcher_AddValidFolder(t *testing.T) {
	w := createWatcher()
	emptyFolder := "../invoices"
	err := w.Add(emptyFolder)
	assert.NoError(t, err)
	assert.NotEmpty(t, w.fileNames)
	assert.NotEmpty(t, w.files)
}
func createWatcher() (*Watcher){
	files := make(map[string]os.FileInfo, 1)
	fileNames := make(map[int]string, 1)
	return &Watcher{files : files, fileNames: fileNames}
}