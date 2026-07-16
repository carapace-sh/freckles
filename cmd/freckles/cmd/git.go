package cmd

import (
	"os"
	"os/exec"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var gitCmd = &cobra.Command{
	Use:                "git",
	Short:              "invoke git on freckles directory",
	DisableFlagParsing: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("git", append([]string{"-C", freckles.Dir()}, args...)...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		return c.Run()
	},
}

// ANCHOR_END: cmd

func init() {
	rootCmd.AddCommand(gitCmd)

	// ANCHOR: positional
	carapace.Gen(gitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin("git", "-C", freckles.Dir())
		}),
	)
	// ANCHOR_END: positional
}
