package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
  "time"
)

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "bump",
	Run: func(cmd *cobra.Command, args []string) {
	  now := time.Now().UTC()

	  formattedTime := now.Format("06.01.02150405")

	  fmt.Println(formattedTime)

	},
}


func init() {
	rootCmd.AddCommand(bumpCmd)
}
