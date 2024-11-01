package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "wwork",
	Short:   "A toolbox for calculations related to woodworking",
	Version: "0.1.0",
}

func init() {
	rootCmd.AddCommand(emcCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
