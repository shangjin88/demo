package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// runCommand runs the cmd and returns the combined stdout and stderr, or an
// error if the command failed.
func runCommand(cmd ...string) (string, error) {
	output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run %q: %s (%s)", strings.Join(cmd, " "), err, output)
	}
	return string(output), nil
}

// getDockerInfo returns the Info struct for the running Docker daemon.
func getDockerInfo() (string, error) {
	output, err := runCommand("docker", "info")
	if err != nil {
		return "", fmt.Errorf("failed to get docker info: %v", err)
	}
	return strings.TrimSpace(output), nil
}

func main() {
	info, _ := getDockerInfo()
	fmt.Println(info)
}
