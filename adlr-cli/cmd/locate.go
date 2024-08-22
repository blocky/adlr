package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/spf13/cobra"
)

var BuildListPath string
var LocatedPath string

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate dependency licenses",
	Long: `Outputs a file containing located licenses. If there was trouble locating
one or more licenses, an error is returned with the list of missing licenses`,
	Run: func(cmd *cobra.Command, args []string) {
		buildlist, err := os.Open(BuildListPath)
		defer buildlist.Close()
		ExitOnErr(err)
		Locate(buildlist)
	},
}

func init() {
	locateCmd.Flags().StringVarP(
		&BuildListPath, "buildlist", "b", "./buildlist.json",
		"Path of module build list in json format",
	)
	locateCmd.Flags().StringVarP(
		&LocatedPath, "located", "l", "./located-licenses.json",
		"Output file containing located licenses",
	)
	licenseCmd.AddCommand(locateCmd)
}

func Locate(buildlist *os.File) {
	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(buildlist)
	ExitOnErr(err)

	missing := ""
	prospects := adlr.MakeProspects(gotool.FilterImportModules(mods)...)
	located, err := adlr.MakeProspector().Prospect(prospects...)
	if err != nil {
		missing = err.Error()
	}

	bytes, err := json.MarshalIndent(located, "", "\t")
	if err != nil {
		fmt.Printf("marshaling prospects: %w", err)
		os.Exit(1)
	}

	err = WriteFile(LocatedPath, bytes)
	if err != nil {
		fmt.Printf("saving found: %w", err)
		os.Exit(1)
	}

	if missing != "" {
		fmt.Printf("%s", missing)
		os.Exit(1)
	}
}
