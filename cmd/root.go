package cmd

import (
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:   "svcli",
	Short: "svcli is a CLI tool for semantic versioning",
	Long:  `svcli helps manage semantic versioning for your projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}
