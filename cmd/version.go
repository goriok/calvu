package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/repos"
  "fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
    v, err := repos.GitRefs.CurrentVersion()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(v.String())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
