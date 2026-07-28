package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list files",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return freckles.Walk(func(freckle freckles.Freckle) error {
			_style := style.ForPathExt(filepath.Join(freckles.Dir(), freckle.Path), carapace.NewContext(args...))
			fmt.Println(format(freckle.Path, _style))
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func format(s, _style string) string {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		return s
	}
	return fmt.Sprintf("\033[%vm%v\033[0m", style.SGR(_style), s)
}
