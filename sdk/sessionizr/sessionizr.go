package sessionizr

import (
	"fmt"
	"iter"

	config "barglvojtech.net/tmux-sessionizr/sdk/config"
	"barglvojtech.net/tmux-sessionizr/sdk/shell"
	"barglvojtech.net/tmux-sessionizr/sdk/tmux"
)

type Sessionizr struct {
	sessions map[string]config.Session
}

func NewSessionizr(sessions map[string]config.Session) *Sessionizr {
	return &Sessionizr{
		sessions: sessions,
	}
}

func (s *Sessionizr) IsValid() bool {
	return len(s.sessions) > 0
}

func (s *Sessionizr) EnsureSession(name string) error {
	session, ok := s.sessions[name]
	if !ok {
		return fmt.Errorf("tmux-sessionizr: session \"%s\" not found", name)
	}

	if tmux.HasSession(session.Name) {
		return nil
	}

	cfg := &tmux.NewSessionConfig{}
	if wd, ok := session.Options[config.OptionWorkingDirectory]; ok {
		cfg.WithWorkingDirectory(wd)
	}

	err := tmux.NewSession(session.Name, cfg)
	if err != nil {
		return fmt.Errorf("tmux-sessionizr: failed to create session %s", session.Name)
	}

	err = tmux.SwitchClient(session.Name)
	if err != nil {
		return fmt.Errorf("tmux-sessionizr: failed to switch to session %s", session.Name)
	}

	for _, shellWords := range session.Commands {
		err = shell.Execute(shellWords...)
		if err != nil {
			return fmt.Errorf("tmux-sessionizr: failed to run command %s", shellWords)
		}
	}

	return nil
}

func (s *Sessionizr) Sessions() iter.Seq[string] {
	return func(yield func(string) bool) {
		for name := range s.sessions {
			if !yield(name) {
				return
			}
		}
	}
}
