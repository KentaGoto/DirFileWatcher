package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// Directory to be monitored.
	dirToWatch := "."

	// Create an instance of fsnotify.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// Display event type and file name.
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				// Display errors.
				fmt.Println("error:", err)
			}
		}
	}()

	// Add directories to be monitored.
	err = watcher.Add(dirToWatch)
	if err != nil {
		log.Fatal(err)
	}
	<-make(chan struct{})
}
