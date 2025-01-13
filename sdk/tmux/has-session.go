package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func HasSession(name string) bool {
	err := shell.Execute("tmux", "has-session", "-t", name)
	return err == nil
}
