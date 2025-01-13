package config

type Session struct {
	Name     string
	Options  map[string]string
	Commands [][]string
}
