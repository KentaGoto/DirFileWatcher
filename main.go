package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Monitoring...")

	// Create a new fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	// Create or open log.log for appending log output
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new logger, writing to log.log
	logger := log.New(file, "", log.LstdFlags)

	// Open SQLite database
	db, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, event TEXT, file TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	// Run a separate goroutine to watch for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event:", event)
				logger.Println("Event:", event)

				// Save event to SQLite database
				statement, err := db.Prepare("INSERT INTO events (event, file) VALUES (?, ?)")
				if err != nil {
					log.Println(err)
					continue
				}
				statement.Exec(event.Op.String(), event.Name)

				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)
					logger.Println("Modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
				logger.Println("Error:", err)
			}
		}
	}()

	// Add multiple directories to the watcher
	directories := []string{
		"/path/to/your/directory1",
		"/path/to/your/directory2",
		// You can add more directories here
	}
	for _, dir := range directories {
		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
	}

	<-done
}
