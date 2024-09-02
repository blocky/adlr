package cmd

import (
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
	type VersionOutput struct {
		Version string `json:"version"`
	}
	p := prettyprinter.NewPrettyPrinter()
	err := p.
		Add(VersionOutput{version}).
		StdoutDump().
		StderrDumpOnError()
	ExitOnErr(err)
}
