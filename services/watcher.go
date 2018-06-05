package services

import (
	"github.com/radovskyb/watcher"
	"fmt"
	"time"
	"github.com/sirupsen/logrus"
)

func WatchFolder(folder string, watcher *watcher.Watcher)  {
	// Watch this folder for changes.
	if err := watcher.Add(folder); err != nil {
		logrus.Fatalln(err)
	}
	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range watcher.WatchedFiles() {
		fmt.Printf("%s: %s\n", path,  f.Name())
	}

	// Trigger 2 events after watcher started.
	go func() {
		watcher.Wait()
	}()

	if err := watcher.Start(time.Millisecond * 100); err != nil {
		logrus.Fatalln(err)
	}
}
