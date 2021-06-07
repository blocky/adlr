package cmd

import (
	_ "embed"

	"github.com/spf13/cobra"

	"github.com/blocky/prettyprinter"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "ADLR version",
	Long:  `Print ADLR semver tag`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion(Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func PrintVersion(
	version string,
) {
	p := prettyprinter.NewPrettyPrinter()
	err := p.
		Add(version).
		StdoutDump().
		StderrDumpOnError()
	ExitOnErr(err)
}
