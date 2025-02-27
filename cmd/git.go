package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/cversion"
  "calvu/internal/git"
  "fmt"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git",
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Help()
	},
}

var gitBumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "bump",
	Run: func(cmd *cobra.Command, args []string) {
    v, err := cversion.FromGit()
    if err != nil {
      fmt.Print(err)
      return
    }
    v.Bump()
    err = git.PushTag(v.Value())
    if err != nil {
      fmt.Print(err)
      return
    }

    fmt.Printf("%s tag pushed\n", v.Value())
	},
}

func init() {
  gitCmd.AddCommand(gitBumpCmd)
	rootCmd.AddCommand(gitCmd)
}
