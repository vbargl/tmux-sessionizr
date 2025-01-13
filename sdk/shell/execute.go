package shell

import "os/exec"

func Execute(shellWords ...string) error {
	cmd := exec.Command(shellWords[0], shellWords[1:]...)
	_, err := cmd.Output()
	return err
}

func ExecuteWithOutput(shellWords ...string) (string, error) {
	cmd := exec.Command(shellWords[0], shellWords[1:]...)
	output, err := cmd.Output()
	return string(output), err
}
