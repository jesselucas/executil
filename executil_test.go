package executil

import (
	"fmt"
	"os"
	"testing"
)

func TestSetOutputChan(t *testing.T) {
	tests := []struct {
		command string
		input   string
		output  string
	}{
		{"echo", "TestSetOutputChan", "TestSetOutputChan"},
		{"echo", "TestSetOutputChan again", "TestSetOutputChan again"},
	}
	for _, test := range tests {
		outputChan := make(chan string)
		cmd := Command(test.command, test.input)
		cmd.OutputChan = outputChan
		err := cmd.StartAndWait()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		out := <-outputChan
		fmt.Println("[chan]" + out)
	}
}

func TestStartAndWait(t *testing.T) {
	tests := []struct {
		command string
		input   string
		output  string
	}{
		{"echo", "TestCmdStart", "TestCmdStart"},
		{"echo", "TestCmdStart again", "TestCmdStart again"},
	}
	for _, test := range tests {
		cmd := Command(test.command, test.input)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func TestSetOutputPrefix(t *testing.T) {
	tests := []struct {
		command string
		prefix  string
		input   string
		output  string
	}{
		{"echo", "new-prefix", "TestSetOutputChan", "TestSetOutputChan"},
		{"echo", "echoandthebunnymen", "TestSetOutputChan again", "TestSetOutputChan again"},
	}
	for _, test := range tests {
		cmd := Command(test.command, test.input)
		cmd.OutputPrefix = test.prefix
		err := cmd.StartAndWait()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
