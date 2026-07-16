package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/freckles/cmd/freckles/cmd/action"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "link dotfiles",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return freckles.Walk(func(freckle freckles.Freckle) error {
			if err := freckle.Symlink(false); err != nil {
				return fmt.Errorf("%s: %w", freckle.Path, err)
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		action.ActionFreckles(),
	)
}
