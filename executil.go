package executil

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// Semantic Version
const VERSION = "0.0.6"

// Cmd struct embeds exec.Cmd
type Cmd struct {
	OutputChan   chan string
	OutputPrefix string
	ShowOutput   bool
	Cmd          *exec.Cmd
}

// Command returns a an executil Cmd struct with
// an exec.Cmd struct embedded in it
func Command(name string, arg ...string) *Cmd {
	c := new(Cmd)
	c.ShowOutput = true
	// set the exec.Cmd
	c.Cmd = exec.Command(name, arg...)
	return c
}

// CombinedOutput wrapper for exec.Cmd.CombinedOutput()
func (c *Cmd) CombinedOutput() ([]byte, error) {
	return c.Cmd.CombinedOutput()
}

// Output wrapper for exec.Cmd.Output()
func (c *Cmd) Output() ([]byte, error) {
	return c.Cmd.Output()
}

// Run wrapper for exec.Cmd.Run()
func (c *Cmd) Run() error {
	return c.Cmd.Run()
}

// Start wrapper for exec.Cmd.Start() with added output if
// ShowOutput bool is true
func (c *Cmd) Start() error {
	// go routines to scan command out and err
	if c.ShowOutput {
		err := c.createPipeScanners()
		if err != nil {
			return err
		}
	}

	return c.Cmd.Start()
}

// StartAndWait starts and waits
func (c *Cmd) StartAndWait() error {
	err := c.Start()
	if err != nil {
		return err
	}

	return c.Wait()
}

// StderrPipe wrapper for exec.Cmd.StderrPipe()
func (c *Cmd) StderrPipe() (io.ReadCloser, error) {
	return c.Cmd.StderrPipe()
}

// StdinPipe wrapper for exec.Cmd.StdinPipe()
func (c *Cmd) StdinPipe() (io.WriteCloser, error) {
	return c.Cmd.StdinPipe()
}

// StdoutPipe wrapper for exec.Cmd.StdoutPipe()
func (c *Cmd) StdoutPipe() (io.ReadCloser, error) {
	return c.Cmd.StdoutPipe()
}

// Wait wrapper for exec.Cmd.Wait()
func (c *Cmd) Wait() error {
	return c.Cmd.Wait()
}

// Create stdout, and stderr pipes for given *Cmd
// Only works with cmd.Start()
func (c *Cmd) createPipeScanners() error {
	// set prefix for pipe scanners
	prefix := c.Cmd.Path
	if c.OutputPrefix != "" {
		prefix = c.OutputPrefix
	}

	stdout, err := c.Cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	stderr, err := c.Cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf(bold("ERROR:")+"\n  Error: %s", err.Error())
	}

	// Created scanners for in, out, and err pipes
	outScanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)

	// Scan for text
	go func() {
		for errScanner.Scan() {
			c.scannerOutput(prefix, errScanner.Text())
		}
	}()

	go func() {
		for outScanner.Scan() {
			c.scannerOutput(prefix, outScanner.Text())
		}
	}()

	return nil
}

func (c *Cmd) scannerOutput(prefix string, text string) {
	out := fmt.Sprintf("[%s] %s\n", prefix, text)
	if c.OutputChan != nil {
		c.OutputChan <- out
	} else {
		fmt.Println(out)
	}
}

func bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
