package services

import (
	"github.com/radovskyb/watcher"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/redis-wrapper"
	"github.com/satori/go.uuid"
)

type WatchedFile struct {
	Name string
	Path string
}

func WatchFolder(folder string, watcher *watcher.Watcher, storage redis_wrapper.Storage) (error){
	// Watch this folder for changes.
	if err := watcher.Add(folder); err != nil {
		return err
	}
	AddFileNameToClient(storage, watcher)

	// Trigger 2 events after watcher started.
	go func() {
		watcher.Wait()
	}()

	if err := watcher.Start(time.Millisecond * 100); err != nil {
		return err
	}
	return nil
}
func AddFileNameToClient(client redis_wrapper.Storage, watcher *watcher.Watcher) {
	files := GetWatchedFiles(watcher)
	for _, file := range files {
		indexName := uuid.NewV4().String()
		err := client.Set(indexName, file.Name, 0)
		if err != nil {
			logrus.Error(err.Error())
		}
	}
}
func GetWatchedFiles(watcher *watcher.Watcher) ([]WatchedFile) {
	files := make([]WatchedFile, 0)
	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range watcher.WatchedFiles() {
		if f.Name() != "" {
			file := WatchedFile{}
			file.Name = f.Name()
			file.Path = path
			files = append(files, file)
		}
	}
	return files
}
