package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
)

func TestWatcher_AddInvalidFolder(t *testing.T) {
	w := NewWatcher("invalid")
	err := w.Add()
	assert.Error(t, err)
}
func TestWatcher_AddValidEmptyFolder(t *testing.T) {
	emptyFolder := "empty"
	w := NewWatcher(emptyFolder)
	os.Mkdir(emptyFolder, os.ModeDir)
	defer os.Remove(emptyFolder)
	err := w.Add()
	assert.Error(t, err)
}
func TestWatcher_AddValidFolder(t *testing.T) {
	folder := "../invoices"
	w := NewWatcher(folder)
	err := w.Add()
	fmt.Printf("%v", w.files)
	fmt.Printf("%v", w.fileNames)
	for _, file := range w.files{
		fmt.Printf("%#v", file.IsDir())
	}
	assert.NoError(t, err)
	assert.NotEmpty(t, w.fileNames)
	assert.NotEmpty(t, w.files)
}
func TestWatcher_StartWithInvalidDuration(t *testing.T) {
	folder := "../invoices"
	w := NewWatcher(folder)
	err := w.Add()
	assert.NoError(t, err)
	err = w.Start(-1)
	assert.Error(t, err)
}
//func TestWatcher_StartWithAlreadyRunnigWatcher(t *testing.T) {
//	folder := "../invoices"
//	w := NewWatcher(folder)
//	err := w.Add()
//	assert.NoError(t, err)
//	err = w.Start(1)
//	w.Start(3)
//	assert.NoError(t, err)
//}