package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/gotool"
	"github.com/blocky/prettyprinter"
)

var BuildListPath string
var ModuleDir string
var Verbose bool

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

	evaluateCmd.Flags().StringVar(
		&BuildListPath, buildListKey, "./buildlist.json", "path of json build list",
	)
	evaluateCmd.Flags().StringVarP(
		&ModuleDir, "dir", "d", "./", "module directory",
	)
	evaluateCmd.Flags().BoolVarP(
		&Verbose, "verbose", "v", false, "verbose debugging output",
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

	prospector := adlr.MakeLicenseProspector()
	mines, err := prospector.Prospect(prospects...)
	ExitOnErr(err)

	miner := adlr.MakeLicenseMiner()
	locks, err := miner.Mine(mines...)
	if Verbose && err != nil {
		PrintStderr(err.Error())
	}

	licenselock := adlr.MakeLicenseLock(ModuleDir)
	err = licenselock.Lock(locks)
	ExitOnErr(err)

	locks, err = licenselock.Read()
	ExitOnErr(err)
	auditor := adlr.MakeLicenseAuditor()
	err = auditor.Audit(locks)
	ExitOnErr(err)
}

func PrintStderr(
	errStr string,
) {
	p := prettyprinter.NewPrettyPrinter()
	err := p.
		Add(errStr).
		StderrDump().
		StderrDumpOnError()
	ExitOnErr(err)
}
