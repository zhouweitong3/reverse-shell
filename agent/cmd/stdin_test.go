package cmd

import (
	"bytes"
	"testing"

	"github.com/maxlaverse/reverse-shell/agent/cmd/test"
	"github.com/stretchr/testify/assert"
)

func TestExecutingCommand(t *testing.T) {
	// Create fake input/output devices
	inputBuf := bytes.NewBufferString("")
	outputBuf := bytes.NewBufferString("")

	// Initialize listener
	listener := newStdinListener(inputBuf, outputBuf)

	// Watch the listener's exit
	doneCh := make(chan error)
	go test.WaitFor(listener.Listen, doneCh)

	// Send test command
	inputBuf.Write([]byte("echo 'THIS IS A TEST';exit\n"))

	// Wait for the listener to exit and test the result
	err := <-doneCh
	assert.Equal(t, err.Error(), "Process terminated")
	assert.Contains(t, outputBuf.String(), "THIS IS A TEST\r\nexit\r\n")
}
