package config

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/junegunn/go-shellwords"
)

var defaultPaths = []string{
	"$HOME/.config/tmux-sessionizr/tmux-sessionizr.sess-sh",
	"$HOME/.config/tmux/conf.d/tmux-sessionizr.sess-sh",
}

func init() {
	for i, path := range defaultPaths {
		defaultPaths[i] = os.ExpandEnv(path)
	}
}

type Loader struct {
	LoadedFiles string
	Sessions    map[string]Session
}

func (l *Loader) LoadPaths(paths ...string) error {
	var (
		sess []Session
		err  error
	)

	switch {
	case len(paths) == 0:
		sess, err = loadPaths(defaultPaths...)
	default:
		sess, err = loadPaths(paths...)
	}

	l.addSessions(sess...)
	return err
}

func (l *Loader) Load(readers ...io.Reader) error {
	var (
		sess []Session
		err  error
	)

	switch {
	case len(readers) == 0:
		sess, err = loadPaths(defaultPaths...)
	default:
		sess, err = loadReaders(readers...)
	}

	l.addSessions(sess...)
	return err
}

func (l *Loader) addSessions(sessions ...Session) {
	if l.Sessions == nil {
		l.Sessions = make(map[string]Session)
	}

	for _, session := range sessions {
		l.Sessions[session.Name] = session
	}
}

func loadPaths(file ...string) ([]Session, error) {
	var (
		sess []Session
		errs []error
	)

	for _, f := range file {
		f, err := os.Open(f)
		if err != nil {
			continue
		}

		s, err := load(f)
		switch {
		case err != nil:
			errs = append(errs, err)
		case err == nil:
			sess = append(sess, s...)
		}

		_ = f.Close()
		break
	}

	return sess, errors.Join(errs...)
}

func loadReaders(r ...io.Reader) ([]Session, error) {
	var (
		sess []Session
		errs []error
	)

	for _, reader := range r {
		if reader == nil {
			continue
		}

		s, err := load(reader)
		switch {
		case err != nil:
			errs = append(errs, err)
		case err == nil:
			sess = append(sess, s...)
		}
		break
	}

	return sess, errors.Join(errs...)
}

func load(r io.Reader) (sess []Session, err error) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	var activeSession *Session
	var activeCommand string

	for sc.Scan() {
		line := sc.Text()

		switch {
		case strings.TrimSpace(line) == "": // empty line
		case strings.HasPrefix(line, "#"): // comment
		case strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]"): // [sessionName]
			parts := strings.Split(strings.Trim(line, "[]"), ":")
			if len(parts) != 2 {
				return nil, errors.New("invalid session name, expected format: [sessionName:workingDirectory]")
			}

			name, wd := parts[0], os.ExpandEnv(parts[1])
			sess = append(sess, Session{Name: name, WorkingDirectory: wd})
			activeSession = &sess[len(sess)-1]
		case strings.HasSuffix(line, "\\"): // multi-line command
			activeCommand += strings.TrimSuffix(line, "\\")
		default: // single-line command
			if activeSession == nil {
				return nil, errors.New("no session name found")
			}

			cmd, err := shellwords.Parse(line)
			if err != nil {
				return nil, err
			}
			activeSession.Commands = append(activeSession.Commands, cmd)
		}
	}

	return sess, nil
}

type Session struct {
	Name             string
	WorkingDirectory string
	Commands         [][]string
}
