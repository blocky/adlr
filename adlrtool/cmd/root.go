package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Run an adlr command",
	Long:  `Run an adlr command, see subcommands for more info.`,
}

func Execute() {
	err := rootCmd.Execute()
	ExitOnErr(err)
}

func ExitOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
