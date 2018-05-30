package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	watchFolder := flag.String("watch_folder", "invoices/", "Name of the folder to watch")
	flag.Parse()
	chan1 := make(chan string, 1)

	for {
		d, _ := os.Open(*watchFolder)
		files, _ := d.Readdir(-1)
		for i, file := range files {
			fmt.Printf("Files in this folder are %s \n", file.Name())
			if i == len(files){
				chan1<-"done"
			}
		}
		close(chan1)
		_, ok := <-chan1
		if !ok{
			break
		}
	}

}
func ReadFolder(folder string) int{
	d, _ := os.Open(folder)
	files, _ := d.Readdir(-1)
	for _, file := range files {
		fmt.Printf("Files in this folder are %s \n", file.Name())

	}
	return len(files)
}
