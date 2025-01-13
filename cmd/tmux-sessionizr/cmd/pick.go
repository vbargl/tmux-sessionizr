package cmd

import (
	"fmt"
	"os"
	"strings"

	"barglvojtech.net/tmux-sessionizr/sdk/randstr"
	"barglvojtech.net/tmux-sessionizr/sdk/tmux"
	"github.com/spf13/cobra"
)

var pickCmd = cobra.Command{
	Use: "pick",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sessionizr := newSessionizr()

		sessions := ""
		for sess := range sessionizr.Sessions() {
			sessions += sess + "\n"
		}

		buffer := "tmux-sessionizr#" + randstr.Hex(8)
		err := tmux.SetBuffer(buffer, sessions)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}

		popupCmd := fmt.Sprintf("tmux show-buffer -b %s | fzf | tmux load-buffer -b %s -", buffer, buffer)
		_ = tmux.DisplayPopup(popupCmd, (&tmux.DisplayPopupConfig{}).WithExitOnFinish(true))

		sess, err := tmux.ShowBuffer(buffer)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}

		sess = strings.TrimSpace(sess)
		_ = tmux.DeleteBuffer(buffer)

		err = sessionizr.EnsureSession(sess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(&pickCmd)
}
