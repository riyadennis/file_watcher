package services

import (
	"github.com/riyadennis/redis-wrapper"
	"time"
	"os"
	"errors"
	"sync"
)

type FolderWatcher struct {
	Folder  string
	Storage redis_wrapper.Storage
	Watcher *Watcher
}
type Watcher struct {
	files     map[string]os.FileInfo
	fileNames map[int]string
	Folder    string
	Event     chan string
	Error     chan error
	Stop      chan bool
	Running   bool
}

func NewFolderWatcher(storage redis_wrapper.Storage, watcher *Watcher, ) *FolderWatcher {
	return &FolderWatcher{
		Storage: storage,
		Watcher: watcher,
	}
}
func NewWatcher(folder string) *Watcher {
	files := make(map[string]os.FileInfo)
	fileNames := make(map[int]string)
	stop := make(chan bool)
	return &Watcher{Folder: folder, files: files, fileNames: fileNames, Stop: stop}
}

func (fw *FolderWatcher) StartWatcher(duration time.Duration) (*Watcher, error) {
	// Watch this folder for changes.
	if err := fw.Watcher.Add(); err != nil {
		return nil, err
	}
	err := fw.Watcher.Start(duration)
	if err != nil {
		return nil, err
	}
	defer fw.Watcher.Close()
	return fw.Watcher, nil
}
func (w Watcher) Add() (error) {
	files, err := getFilesFromFolder(w.Folder)
	if err != nil {
		return err
	}
	if len(files) < 0 {
		return errors.New("Empty folder")
	}
	for k, _ := range files {
		i := 0
		w.fileNames[i] = k
		i++
	}
	w.files = files
	return nil
}
func getFilesFromFolder(folder string) (map[string]os.FileInfo, error) {
	returnFiles := make(map[string]os.FileInfo)
	_, err := os.Stat(folder)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(folder)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	//read all the contents
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.Name() != ""{
			returnFiles[f.Name()] = f
		}
	}
	return returnFiles, nil
}
func (w Watcher) Start(duration time.Duration) (error) {
	if duration < time.Nanosecond {
		return errors.New("invalid duration")
	}

	if w.Running == true {
		w.Error <- errors.New("Watcher already running")
	}
	w.Running = true
	wt := sync.WaitGroup{}
	wt.Wait()
	for {
		//get current files
		files, err := getFilesFromFolder(w.Folder)
		if err != nil {
			w.Error <- err
		}
		go func() {
			w.watchEvents(files)
		}()

	}
	return nil
}

func (w Watcher) watchEvents(files map[string]os.FileInfo) {
	for path, currentFile := range files {
		oldFile := w.files[path]
		if w.files[path] != nil {
			if currentFile.ModTime() != oldFile.ModTime() {
				w.Event <- "file changed"
			}
		}
	}

}
func (w Watcher) Close() {

}
