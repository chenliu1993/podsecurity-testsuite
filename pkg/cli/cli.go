package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func ExecCommand(cmdName string, input io.Reader, args ...string) (string, error) {
	var output bytes.Buffer
	writer := io.MultiWriter(&output, os.Stderr)

	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = input
	cmd.Stdout = writer
	cmd.Stderr = writer
	if err := cmd.Run(); err != nil {
		returnCode := -1
		if exitError, ok := err.(*exec.ExitError); ok {
			returnCode = exitError.ExitCode()
		}
		return output.String(), fmt.Errorf("command failed with code: %v", returnCode)
	}
	return output.String(), nil
}

func Kubectl(input io.Reader, args ...string) (string, error) {
	return ExecCommand("kubectl", input, args...)
}
