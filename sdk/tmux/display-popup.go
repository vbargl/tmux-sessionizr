package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func DisplayPopup(cmd string, config *DisplayPopupConfig) error {
	shellWords := []string{"tmux", "display-popup"}
	if config.ExitOnFinish {
		shellWords = append(shellWords, "-E")
	}
	shellWords = append(shellWords, cmd)
	return shell.Execute(shellWords...)
}

type DisplayPopupConfig struct {
	ExitOnFinish bool
}

func (c *DisplayPopupConfig) WithExitOnFinish(exitOnFinish bool) *DisplayPopupConfig {
	c.ExitOnFinish = exitOnFinish
	return c
}
