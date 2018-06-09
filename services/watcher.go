package services

import (
	"github.com/radovskyb/watcher"
	"time"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/redis-wrapper"
	"github.com/satori/go.uuid"
	"os"
	"github.com/pkg/errors"
)

type WatchedFile struct {
	Name string
	Path string
	Event chan struct{}
	Error chan error
}
type Watcher struct {
	files     map[string]os.FileInfo
	fileNames map[int]string
}

func (w Watcher) Add(folder string) (error) {
	files, err := getFilesFromFolder(folder)
	if err != nil {
		return err
	}
	if len(files) < 0 {
		return errors.New("Empty folder")
	}
	for k, f := range files {
		w.fileNames[k] = f.Name()
		w.files[f.Name()] = f
	}
	return nil
}
func getFilesFromFolder(folder string) ([]os.FileInfo, error) {
	_, err := os.Stat(folder)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	//read all the contents
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	return files, nil
}
func (w Watcher) Start(duration time.Duration) (error) {
	return nil
}
func (w Watcher) Close() {

}

func CreateWatcher() (*watcher.Watcher) {
	watcher := watcher.New()
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				logrus.Println("Event:", ev)
			case err := <-watcher.Error:
				logrus.Println("Error:", err)
			}
		}
	}()
	return watcher
}
func CreateRedisClient() (*redis_wrapper.Client) {
	client := &redis_wrapper.Client{}
	redisClient, err := client.Create()
	if err != nil {
		logrus.Errorf("Unable to connect to the storage got error :  %s", err)
		return nil
	}
	return redisClient
}
func WatchFolder(folder string, watcher *watcher.Watcher, storage redis_wrapper.Storage) (error) {
	// Watch this folder for changes.
	if err := watcher.Add(folder); err != nil {
		return err
	}
	files := GetWatchedFiles(watcher)
	AddFilesToStorage(files, storage)

	// Trigger 2 events after watcher started.
	go func() {
		watcher.Wait()
	}()

	if err := watcher.Start(time.Millisecond * 100); err != nil {
		return err
	}
	return nil
}
func AddFilesToStorage(files []WatchedFile, client redis_wrapper.Storage) {
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
