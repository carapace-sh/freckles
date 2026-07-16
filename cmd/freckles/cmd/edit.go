package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/freckles/cmd/freckles/cmd/action"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var editCmd = &cobra.Command{
	Use:   "edit [FILE]",
	Short: "edit a dotfile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command(editor(), filepath.Join(freckles.Dir(), args[0]))
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

// ANCHOR_END: cmd

func init() {
	rootCmd.AddCommand(editCmd)

	// ANCHOR: positional
	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionFreckles(),
	)
	// ANCHOR_END: positional
}

func editor() string { // source github.com/cli/cli
	defaultEditor := "vi"
	if runtime.GOOS == "windows" {
		defaultEditor = "notepad"
	} else if value := os.Getenv("GIT_EDITOR"); value != "" {
		defaultEditor = value
	} else if value := os.Getenv("VISUAL"); value != "" {
		defaultEditor = value
	} else if value := os.Getenv("EDITOR"); value != "" {
		defaultEditor = value
	}
	return defaultEditor
}
