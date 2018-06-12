package main

import (
	"flag"
	"github.com/file_watcher/services"
	"github.com/sirupsen/logrus"
	"fmt"
)

func main() {
	watchFolder := flag.String("watch_folder", "invoices/", "Name of the folder to watch")
	watchType := flag.String("type", "radovskyb", "Type of watch")
	flag.Parse()
	watcher := services.CreateWatcher()
	redisClient := services.CreateRedisClient()
	if redisClient == nil {
		return
	}
	if *watchType == "radovskyb"{
		err := services.WatchFolder(*watchFolder, watcher, redisClient)
		if err != nil {
			logrus.Error(err)
		}
		return
	}

	w := services.NewWatcher(*watchFolder)
	w.Add()
	w.Start(1)
	//watcher.Close()
	go func(){
		for{
			select{
				case event := <- w.Event:
					fmt.Printf("Event fired %s", event)
			}
		}
	}()
}
