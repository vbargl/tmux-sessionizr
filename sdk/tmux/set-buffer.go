package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func SetBuffer(buffer string, content string) error {
	return shell.Execute("tmux", "set-buffer", "-b", buffer, content)
}
