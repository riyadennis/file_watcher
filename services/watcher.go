package services

import (
	"github.com/radovskyb/watcher"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/go-redis/redis"
	"github.com/riyadennis/redis-wrapper"
	"github.com/satori/go.uuid"
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
	c := redis_wrapper.Client{}
	client, err := c.Create()
	if err != nil {
		logrus.Error(err.Error())
	}

	AddFileNameToClient(client.RedisClient, watcher)

	// Trigger 2 events after watcher started.
	go func() {
		watcher.Wait()
	}()

	if err := watcher.Start(time.Millisecond * 100); err != nil {
		logrus.Fatalln(err)
	}
}
func AddFileNameToClient(client *redis.Client, watcher *watcher.Watcher) {
	files := GetWatchedFiles(watcher)
	for _, file := range files {
		indexName := uuid.NewV4().String()
		err := client.Set(indexName, file.Path, 0).Err()
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
