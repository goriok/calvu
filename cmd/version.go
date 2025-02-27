package cmd

import (
	"github.com/spf13/cobra"
  "calvu/internal/version"
  "calvu/internal/git"
  "fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Run: func(cmd *cobra.Command, args []string) {
    v, err := git.HeadDate()
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(version.FromTime(*v))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
