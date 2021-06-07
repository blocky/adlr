package cmd

import "github.com/spf13/cobra"

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Software information",
	Long:  `Display general software information`,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
