package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/version"
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
    headTime, err := git.HeadDate()
    if err != nil{
      fmt.Println(err)
      return
    }

    v := version.FromTime(*headTime)
    err = git.Bump(v)
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
