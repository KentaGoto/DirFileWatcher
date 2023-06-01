# DirFileWatcher

This repository contains a simple Go application that monitors changes in multiple directories using the fsnotify library. Monitored events are output to the console and simultaneously logged to a file called log.log and to the SQLite database events.db.

## Getting Started

### Prerequisites

- Go (at least version 1.16)
- [fsnotify](https://github.com/fsnotify/fsnotify)

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/KentaGoto/DirFileWatcher
```

Navigate to the project directory:

```bash
cd DirFileWatcher
```

## Usage

1. Open `main.go` and modify the `directories` slice with the paths to the directories you want to monitor:

```go
directories := []string{
    "/path/to/your/directory1",
    "/path/to/your/directory2",
    // You can add more directories here
}
```

2. Run the program:

```bash
go run main.go
```

3. Monitored events are logged to the console, to the log.log file, and to the SQLite database.

## Database
The application uses a SQLite database named `events.db`. This database contains a table named `events`, where each event is stored as a new record. The table has the following three columns:

- `id`: A unique ID for the event. This is an integer that auto-increments.
- `event`: The type of event that occurred, such as file creation, modification, or deletion.
- `file`: The name of the file where the event occurred.

## Log File
The `log.log` file contains logs of all file change events. Whenever a new event occurs, its information is appended to this file.

## License

MIT

## Author

Kenta Goto
