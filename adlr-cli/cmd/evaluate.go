package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/blocky/prettyprinter"
)

var BuildListPath string
var ModuleDir string
var Verbose bool
var LicenseList []string

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Evaluate module dependencies",
	Long:  `A command for evaluating module dependency licensing.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildlist, err := os.Open(BuildListPath)
		defer buildlist.Close()
		ExitOnErr(err)

		Evaluate(buildlist)
	},
}

func init() {
	buildListKey := "buildlist"

	evaluateCmd.Flags().StringVarP(
		&BuildListPath, buildListKey, "b", "./buildlist.json",
		"path of module build list in json format",
	)
	evaluateCmd.Flags().StringVarP(
		&ModuleDir, "dir", "d", "./",
		"output directory or location of your existing license.lock",
	)
	evaluateCmd.Flags().BoolVarP(
		&Verbose, "verbose", "v", false,
		"verbose debugging output",
	)
	evaluateCmd.Flags().StringSliceVarP(
		&LicenseList, "whitelist", "w", adlr.DefaultWhitelist,
		"comma separated list of acceptable licenses in SPDX License Identifier format",
	)

	rootCmd.AddCommand(evaluateCmd)
	evaluateCmd.MarkFlagRequired(buildListKey)
}

func Evaluate(buildlist *os.File) {
	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(buildlist)
	ExitOnErr(err)

	direct := gotool.FilterDirectImportModules(mods)
	prospects := adlr.MakeProspects(direct...)

	prospector := adlr.MakeProspector()
	mines, err := prospector.Prospect(prospects...)
	ExitOnErr(err)

	miner := adlr.MakeMiner()
	locks, err := miner.Mine(mines...)
	if Verbose && err != nil {
		PrintStderr(err)
	}

	licenselock := adlr.MakeLicenseLockManager(ModuleDir)
	err = licenselock.Lock(locks...)
	ExitOnErr(err)

	locks, err = licenselock.Read()
	ExitOnErr(err)

	whitelist := adlr.MakeWhitelist(LicenseList)
	auditor := adlr.MakeAuditor(whitelist)
	err = auditor.Audit(locks...)
	ExitOnErr(err)
}

func PrintStderr(
	err error,
) {
	kve := prettyprinter.MakeKeyValueError(err)
	p := prettyprinter.NewPrettyPrinter()
	err = p.
		Add(kve).
		StderrDump().
		StderrDumpOnError()
	ExitOnErr(err)
}
