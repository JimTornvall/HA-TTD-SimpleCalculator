package main

import (
	"os/exec"
	"testing"
)

// untested unfinished code for future reference
func TestReadInput(t *testing.T) {
	cmd := exec.Command("./myapp") // Replace with your actual binary name

	// Create pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("Error creating stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		t.Fatalf("Error creating stderr pipe: %v", err)
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		t.Fatalf("Error starting process: %v", err)
	}

	// Read stdout and stderr
	stdoutOutput := make(chan string)
	stderrOutput := make(chan string)
	go func() {
		readOutput(stdoutPipe, stdoutOutput)
	}()
	go func() {
		readOutput(stderrPipe, stderrOutput)
	}()

	// Wait for the process to finish
	if err := cmd.Wait(); err != nil {
		t.Fatalf("Error waiting for process: %v", err)
	}

	// Print the captured output
	t.Logf("Standard Output:\n%s", <-stdoutOutput)
	t.Logf("Standard Error:\n%s", <-stderrOutput)
}

func readOutput(pipe io.Reader, output chan<- string) {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(pipe)
	output <- buf.String()
}
