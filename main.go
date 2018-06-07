package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/file_watcher/services"
)

func main() {
	watchFolder := flag.String("watch_folder", "invoices/", "Name of the folder to watch")
	flag.Parse()
	watcher := services.CreateWatcher()
	redisClient := services.CreateRedisClient()
	if redisClient == nil {
		return
	}
	err := services.WatchFolder(*watchFolder, watcher, redisClient)
	if err != nil {
		logrus.Error(err)
	}

	watcher.Close()
}
