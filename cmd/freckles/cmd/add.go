package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var addCmd = &cobra.Command{
	Use:   "add [FILE]...",
	Short: "add dotfiles",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, arg := range args {
			freckle := freckles.Freckle{Path: arg}
			if err := freckle.Add(false); err != nil {
				return fmt.Errorf("%s: %w", arg, err)
			}
		}
		return nil
	},
}

// ANCHOR_END: cmd

func init() {
	rootCmd.AddCommand(addCmd)

	// ANCHOR: positional
	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			batch := carapace.Batch(
				carapace.ActionFiles(),
			)
			if c.Value == "" {
				batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					c.Value = "."
					return carapace.ActionFiles().Invoke(c).ToA()
				}))
			}
			return batch.ToA().ChdirF(traverse.UserHomeDir)
		}),
	)
	// ANCHOR: positional
}
