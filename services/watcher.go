package services

import (
	"github.com/radovskyb/watcher"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
)

type WatchedFile struct {
	Name string
	Path string
}

func WatchFolder(folder string, watcher *watcher.Watcher) {
	// Watch this folder for changes.
	if err := watcher.Add(folder); err != nil {
		logrus.Fatalln(err)
	}
	// Print a list of all of the files and folders currently
	// being watched and their paths.
	files := GetWatchedFiles(watcher)
	for _, file := range files {
		fmt.Printf("File Name: %s \n", file.Name)
	}
	// Trigger 2 events after watcher started.
	go func() {
		watcher.Wait()
	}()

	if err := watcher.Start(time.Millisecond * 100); err != nil {
		logrus.Fatalln(err)
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
