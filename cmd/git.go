package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/repos"
  "calvu/internal/git"
  "calvu/internal/calver"
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
    cver, err := repos.GitRefs.CurrentVersion()
    if err != nil {
      fmt.Print(err)
      return
    }

    v := calver.Bump(*cver)
    err = git.PushTag(v.String())
    if err != nil {
      fmt.Print(err)
      return
    }

    fmt.Printf("%s tag pushed\n", v)
	},
}

func init() {
  gitCmd.AddCommand(gitBumpCmd)
	rootCmd.AddCommand(gitCmd)
}
