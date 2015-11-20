package executil

import (
	"bufio"
	"fmt"
	"os/exec"
)

// Semantic Version
const VERSION = "0.0.4"

// Cmd struct embeds exec.Cmd
type Cmd struct {
	OutputChan   chan string
	OutputPrefix string
	*exec.Cmd
}

// Command returns a an executil Cmd struct with
// an exec.Cmd struct embedded in it
func Command(name string, arg ...string) *Cmd {
	cmd := new(Cmd)

	// set the exec.Cmd
	cmd.Cmd = exec.Command(name, arg...)
	return cmd
}

// SetOutputChan setter function to set OutputChan
func (cmd *Cmd) SetOutputChan(outputChan chan string) {
	cmd.OutputChan = outputChan
}

// SetOutputPrefix setter function to set OutputPrefix
func (cmd *Cmd) SetOutputPrefix(prefix string) {
	cmd.OutputPrefix = prefix
}

// StartWithOutput creates exec.Command and calls Start()
func (cmd *Cmd) StartWithOutput() error {
	// set prefix for pipe scanners
	prefix := cmd.Cmd.Path
	if cmd.OutputPrefix != "" {
		prefix = cmd.OutputPrefix
	}

	// go routines to scan command out and err
	err := createPipeScanners(cmd, prefix)
	if err != nil {
		return err
	}

	// start the exec.Cmd
	return start(cmd.Cmd)
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
func createPipeScanners(cmd *Cmd, prefix string) error {
	stdout, err := cmd.Cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	stderr, err := cmd.Cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	// Created scanners for in, out, and err pipes
	outScanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)

	// Scan for text
	go func() {
		for errScanner.Scan() {
			cmd.scannerOutput(prefix, errScanner.Text())
		}
	}()

	go func() {
		for outScanner.Scan() {
			cmd.scannerOutput(prefix, outScanner.Text())
		}
	}()

	return nil
}

func (cmd *Cmd) scannerOutput(prefix string, text string) {
	out := fmt.Sprintf("[%s] %s\n", prefix, text)
	if cmd.OutputChan != nil {
		cmd.OutputChan <- out
	} else {
		fmt.Println(out)
	}
}

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
