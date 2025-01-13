package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

// SetBuffer sets the content of a tmux buffer.
func ShowBuffer(buffer string) (string, error) {
	return shell.ExecuteWithOutput("tmux", "show-buffer", "-b", buffer)
}
