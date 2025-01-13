package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func SwitchClient(target string) error {
	return shell.Execute("tmux", "switch-client", "-t", target)
}
