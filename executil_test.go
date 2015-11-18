package executil

import (
	"fmt"
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
		CmdStart(test.command, test.input)

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
		CmdStart(test.command, test.input)
	}
}
