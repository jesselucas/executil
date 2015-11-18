package executil

import (
	"bufio"
	"fmt"
	"os/exec"
)

// If OutputChan is set the stdout
// and stderr will be sent to it
var OutputChan chan string

// If OutputPrefix is not set the
// command will be the prefix
var OutputPrefix string

// Semantic Version
const VERSION = "0.0.3"

// SetOutputChan setter function to set OutputChan
func SetOutputChan(outputChan chan string) {
	OutputChan = outputChan
}

// SetOutputPrefix setter function to set OutputPrefix
func SetOutputPrefix(prefix string) {
	OutputPrefix = prefix
}

// CmdStart creates exec.Command and calls Start()
func CmdStart(commandName string, arg ...string) error {
	// run protoc command (protoc --go_out=plugins=grpc:. $proto)
	// execute cmd
	cmd := exec.Command(commandName, arg...)

	// set prefix for pipe scanners
	prefix := commandName
	if OutputPrefix != "" {
		prefix = OutputPrefix
	}

	// go routines to scan command out and err
	err := createPipeScanners(cmd, prefix)
	if err != nil {
		return err
	}

	// start the command
	return start(cmd)
}

func start(cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	return nil
}

// Create stdout, and stderr pipes for given *Cmd
// Only works with cmd.Start()
func createPipeScanners(cmd *exec.Cmd, prefix string) error {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

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

	return nil
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
