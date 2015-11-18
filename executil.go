package executil

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// If OutputChan is set the stdout
// and stderr will be sent to it
var OutputChan chan string

// SetOutputChan Setter function to set OutputChan
func SetOutputChan(outputChan chan string) {
	OutputChan = outputChan
}

// CmdStart creates exec.Command and calls Start()
func CmdStart(commandName string, arg ...string) {
	// run protoc command (protoc --go_out=plugins=grpc:. $proto)
	// execute cmd
	cmd := exec.Command(commandName, arg...)

	// go routines to scan command out and err
	createPipeScanners(cmd, commandName)

	// start the command
	start(cmd)
}

func start(cmd *exec.Cmd) {
	if err := cmd.Start(); err != nil {
		fmt.Printf(bold("ERROR:")+"\n  Error: %s", err.Error())
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Printf(bold("ERROR:")+"\n  Error: %s", err.Error())
		os.Exit(1)
	}
}

// Create stdout, and stderr pipes for given *Cmd
// Only works with cmd.Start()
func createPipeScanners(cmd *exec.Cmd, prefix string) {
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	// Created scanners for in, out, and err pipes
	outScanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)

	// Scan for text
	go func() {
		for errScanner.Scan() {
			scannerOutput(prefix, errScanner.Text())
		}
	}()

	go func() {
		for outScanner.Scan() {
			scannerOutput(prefix, outScanner.Text())
		}
	}()
}

func scannerOutput(prefix string, text string) {
	out := fmt.Sprintf("[%s] %s\n", prefix, text)
	if OutputChan != nil {
		OutputChan <- out
	} else {
		fmt.Println(out)
	}
}

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
