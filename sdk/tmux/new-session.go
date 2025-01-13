package tmux

import "barglvojtech.net/tmux-sessionizr/sdk/shell"

func NewSession(sessionName string, cfg *NewSessionConfig) error {
	cmd := []string{"tmux", "new-session", "-d", "-s", sessionName}
	if cfg.workdir != "" {
		cmd = append(cmd, "-c", cfg.workdir)
	}
	if cfg.env != nil {
		for _, e := range cfg.env {
			cmd = append(cmd, "-e", e)
		}
	}

	return shell.Execute(cmd...)
}

type NewSessionConfig struct {
	workdir string
	env     []string
}

func (cfg *NewSessionConfig) WithWorkingDirectory(workdir string) *NewSessionConfig {
	cfg.workdir = workdir
	return cfg
}

func (cfg *NewSessionConfig) WithEnvironment(key, value string) *NewSessionConfig {
	cfg.env = append(cfg.env, key+"="+value)
	return cfg
}
