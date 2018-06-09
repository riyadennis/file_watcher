package services

import (
	"github.com/riyadennis/redis-wrapper"
	"time"
)

type FolderWatcher struct {
	Folder  string
	Storage redis_wrapper.Storage
	Watcher *Watcher
}

func NewFolderWatcher(storage redis_wrapper.Storage, watcher *Watcher,) *FolderWatcher {
	return &FolderWatcher{
		Storage: storage,
		Watcher: watcher,
	}
}
func (fw *FolderWatcher) StartWatcher(folder string, duration time.Duration) (*Watcher, error){
	// Watch this folder for changes.
	if err := fw.Watcher.Add(folder); err != nil{
		return nil, err
	}
	err := fw.Watcher.Start(duration)
	if err != nil {
		return nil, err
	}
	defer fw.Watcher.Close()
	return fw.Watcher, nil
}
