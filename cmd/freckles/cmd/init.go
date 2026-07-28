package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	spec "github.com/carapace-sh/carapace-spec"
	"github.com/carapace-sh/freckles/pkg/freckles"
	"github.com/spf13/cobra"
)

// ANCHOR: cmd
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init freckles folder",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("git", "init", freckles.Dir())
		if cmd.Flag("clone").Changed {
			c = exec.Command("git", "clone", cmd.Flag("clone").Value.String(), freckles.Dir())
		}
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return err
		}

		if _, err := os.Stat(filepath.Join(freckles.Dir(), ".frecklesignore")); os.IsNotExist(err) {
			return os.WriteFile(filepath.Join(freckles.Dir(), ".frecklesignore"), []byte(".git\n.frecklesignore\n"), os.ModePerm)
		}
		return nil
	},
}

// ANCHOR_END: cmd

func init() {
	initCmd.Flags().String("clone", "", "clone existing repo")

	rootCmd.AddCommand(initCmd)

	// ANCHOR: flagcompletion
	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"clone": spec.ActionMacro("$carapace.tools.git.RepositorySearch"),
	})
	// ANCHOR_END: flagcompletion
}
