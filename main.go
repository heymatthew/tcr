package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func main() {
	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for directories
	err := filepath.WalkDir(".", watchDir)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	done := make(chan bool)

	testCommand := flag.String("c", "go test", "test command")
	flag.Parse()

	// https://medium.com/@skdomino/watch-this-file-watching-in-go-5b5a247cf71f
	go func() {
		for {
			output := make(chan fsnotify.Event, 1)
			go debounce(500*time.Millisecond, watcher.Events, output)

			select {
			case event := <-output:
				fmt.Printf(" - %s - ", event)
				if strings.HasPrefix(event.Name, ".git") {
					continue
				}
				err := run(*testCommand)
				if err != nil {
					run("git checkout .")
				} else {
					run("git add .")
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func debounce(interval time.Duration, in chan fsnotify.Event, out chan fsnotify.Event) {
	var item fsnotify.Event
	timer := time.NewTimer(interval)
	for {
		select {
		case item = <-in:
			timer.Reset(interval)
		case <-timer.C:
			out <- item
		}
	}
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, d fs.DirEntry, err error) error {
	if err != nil {
		fmt.Printf("Error walking %s: %s\n", path, err)
		return err
	}

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if d.IsDir() {
		watcher.Add(path)
	}

	return nil
}

func run(cmd string) error {
	fmt.Printf("\n> %s\n", cmd)
	runnable := strings.Split(cmd, " ")
	runner := exec.Command(runnable[0], runnable[1:]...)
	runner.Stdout = os.Stdout
	return runner.Run()
}
