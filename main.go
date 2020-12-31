package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"
)

func main() {
	testCommand := flag.String("c", "go test", "test command")
	flag.Parse()
	err := run(*testCommand)
	if err != nil {
		run("git checkout .")
	} else {
		run("git add .")
	}
}

func run(cmd string) error {
	runnable := strings.Split(cmd, " ")
	runner := exec.Command(runnable[0], runnable[1:]...)
	runner.Stdout = os.Stdout
	return runner.Run()
}
