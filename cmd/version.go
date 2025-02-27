package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/cversion"
  "fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
    v, err := cversion.FromGit()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(v.Value())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
