package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/cversion"
  "calvu/internal/helm"
  "calvu/internal/git"
  "fmt"
)

var commit bool
var push bool

var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "helm",
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Help()
	},
}

var helmBumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "bump",
	Run: func(cmd *cobra.Command, args []string) {
    v, err := cversion.FromGit()
    if err != nil {
      fmt.Print(err)
      return
    }
    v.Bump()
    chart, err := helm.SetVersion(v.Value())
    if err != nil {
      fmt.Print(err)
      return
    }

    if !commit {
      return
    }

    msg := fmt.Sprintf("chore: version chart %s updated, version: %s",chart.Name, chart.Version)
    err = git.Commit(helm.ChartFile, msg)
    if err != nil {
      fmt.Print(err)
      return
    }

    if !push {
      return
    }

    err = git.Push()
    if err != nil {
      fmt.Print(err)
      return
    }

    err = git.PushTag(v.Value())
    if err != nil {
      fmt.Print(err)
      return
    }
	},
}

func init() {
	helmBumpCmd.Flags().BoolVar(&commit, "commit", false, "Commit changes after bumping version")
	helmBumpCmd.Flags().BoolVar(&push, "push", false, "Push changes after commit changes")

  helmCmd.AddCommand(helmBumpCmd)
	rootCmd.AddCommand(helmCmd)
}
