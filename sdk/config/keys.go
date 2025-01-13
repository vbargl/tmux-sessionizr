package config

import "strings"

const (
	OptionWorkingDirectory = "working-directory"
	OptionSticky           = "sticky"
)

func normalizeOptionKey(key string) string {
	switch key := strings.ToLower(key); key {
	case "wd", "workdir", "workingdir", OptionWorkingDirectory:
		return OptionWorkingDirectory
	case OptionSticky:
		return OptionSticky
	default:
		return key
	}
}
