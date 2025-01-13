package cmd

import (
	"fmt"
	"os"

	"barglvojtech.net/tmux-sessionizr/sdk/config"
	"barglvojtech.net/tmux-sessionizr/sdk/sessionizr"
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use: "tmux-sessionizr",
	
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newSessionizr() *sessionizr.Sessionizr {
	var loader config.Loader
	if err := loader.Load(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	return sessionizr.NewSessionizr(loader.Sessions)
}
