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
		SetOutputChan(outputChan)
		err := CmdStart(test.command, test.input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		out := <-outputChan
		fmt.Println(out)
	}
}

func TestCmdStart(t *testing.T) {
	tests := []struct {
		command string
		input   string
		output  string
	}{
		{"echo", "TestCmdStart", "TestCmdStart"},
		{"echo", "TestCmdStart again", "TestCmdStart again"},
	}
	for _, test := range tests {
		SetOutputChan(nil) // make sure there isn't an outputChan set
		err := CmdStart(test.command, test.input)
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
		SetOutputChan(nil) // make sure there isn't an outputChan set
		SetOutputPrefix(test.prefix)
		err := CmdStart(test.command, test.input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
