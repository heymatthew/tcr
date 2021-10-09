package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var clearScreen = "\033c"

func main() {
	testCommand := flag.String("c", "go test", "test command")
	flag.Parse()
	fmt.Printf(clearScreen)

	err := run(*testCommand)
	if err != nil {
		run("git checkout .")
	} else {
		run("git add .")
	}
}

func run(cmd string) error {
	fmt.Printf("\n> %s\n", cmd)
	runnable := strings.Split(cmd, " ")
	runner := exec.Command(runnable[0], runnable[1:]...)
	runner.Stdout = os.Stdout
	return runner.Run()
}
