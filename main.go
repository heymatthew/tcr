package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	testCommand := flag.String("c", "go test", "test command")
	flag.Parse()
	fmt.Println("Test command:", *testCommand)
	err := run("git", "status")
	if err != nil {
		panic(err)
	}
}

func run(runnable string, args ...string) error {
	cmd := exec.Command(runnable, args...)
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
