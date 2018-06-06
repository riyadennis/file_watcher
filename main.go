package main

import (
	"flag"
	"github.com/radovskyb/watcher"
	"github.com/sirupsen/logrus"
	"github.com/file_watcher/services"
	"github.com/riyadennis/redis-wrapper"
)

func main() {
	watchFolder := flag.String("watch_folder", "invoices/", "Name of the folder to watch")
	flag.Parse()
	watcher := watcher.New()
	done := make(chan bool)
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

	client := &redis_wrapper.Client{}
	redisClient, err := client.Create()
	if err != nil {
		logrus.Error(err)
	}

	err = services.WatchFolder(*watchFolder, watcher, redisClient)
	if err != nil {
		logrus.Error(err)
	}
	<-done
	watcher.Close()
}
