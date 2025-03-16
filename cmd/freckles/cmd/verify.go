package cmd

import (
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/rsteube/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify symlink status",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		freckles.Walk(func(freckle freckles.Freckle) error {
			_style := style.ForPathExt(freckles.Dir()+"/"+freckle.Path, carapace.NewContext(args...))
			if freckle.Verify() {
				fmt.Printf("[%v] %v\n", format("OK", style.Green), format(freckle.Path, _style))
			} else {
				fmt.Printf("[%v] %v\n", format("ERR", style.Red), format(freckle.Path, _style))
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}
