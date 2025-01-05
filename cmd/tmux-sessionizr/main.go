package main

import (
	"fmt"
	"os"
	"os/exec"

	"barglvojtech.net/tmux-sessionizr/sdk/config"
)

func main() {
	tmuxMode()

}

func tmuxMode() {
	var loader config.Loader
	if err := loader.Load(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	for _, session := range loader.Sessions {
		if shell("tmux", "has-session", "-t", session.Name) == nil {
			continue
		}

		err := shell("tmux", "new-session", "-d", "-c", session.WorkingDirectory, "-s", session.Name)
		if err != nil {
			_ = shell("tmux", "display-message", fmt.Sprintf("tmux-sessionizr: failed to create session %s", session.Name))
			os.Exit(1)
		}

		for _, command := range session.Commands {
			err := shell(command...)
			if err != nil {
				_ = shell("tmux", "display-message", fmt.Sprintf("tmux-sessionizr: failed to run command %s", command))
				os.Exit(1)
			}
		}
	}
}

func shell(command ...string) error {
	cmd := exec.Command(command[0], command[1:]...)
	_, err := cmd.Output()
	return err
}
