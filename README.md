# DirFileWatcher

This repository contains a simple Go application that uses the fsnotify library to monitor changes in multiple directories. It outputs the events to the console and also logs them to a file named `log.log`.

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

3. You should see logs printed to the console and logged to a file named `log.log` whenever an event occurs in any of the directories being monitored.

## License

MIT

## Author

Kenta Goto
