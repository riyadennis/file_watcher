package main

import (
	"flag"
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"time"
)

func main() {
	watchFolder := flag.String("watch_folder", "invoices/", "Name of the folder to watch")
	flag.Parse()
	fmt.Printf("%v", *watchFolder)
	watcher := watcher.New()
	done := make(chan bool)
	go func(){
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("Event:", ev)
			case err := <- watcher.Error:
				log.Println("Error:", err)
			}
		}
	}()

	WatchFolder(*watchFolder, watcher)

	<-done
	watcher.Close()
}
func WatchFolder(folder string, watcher *watcher.Watcher)  {
	// Watch this folder for changes.
	if err := watcher.Add(folder); err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}
}
