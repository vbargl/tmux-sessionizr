package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func DisplayMessage(message string) error {
	return shell.Execute("tmux", "display-message", message)
}
