package tmux

import (
	"barglvojtech.net/tmux-sessionizr/sdk/shell"
)

func DeleteBuffer(buffer string) error {
	return shell.Execute("tmux", "delete-buffer", "-b", buffer)
}
