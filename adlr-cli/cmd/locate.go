package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/spf13/cobra"
)

var BuildlistFile string
var LocatedFile string
var ExemptMods []string

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate dependency licenses",
	Long:  "Outputs a json file containing located licenses.",
	Run: func(cmd *cobra.Command, args []string) {
		err := Locate(BuildlistFile, LocatedFile)
		ExitOnErr(err)
	},
}

func init() {
	locateCmd.Flags().StringVarP(
		&BuildlistFile, "buildlist", "b", "./buildlist.json",
		"Path of module build list in json format",
	)
	locateCmd.Flags().StringVarP(
		&LocatedFile, "located", "l", "./located-licenses.json",
		"Output file containing located licenses",
	)
	locateCmd.Flags().StringSliceVarP(
		&ExemptMods, "exempt-modules", "e", []string{},
		"Comma separated list of modules to ignore during license location")
	licenseCmd.AddCommand(locateCmd)
}

func Locate(
	buildlistFile string,
	locatedFile string,
) error {
	buildlist, err := os.ReadFile(buildlistFile)
	if err != nil {
		return fmt.Errorf("reading buildlist file: %w", err)
	}

	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(bytes.NewReader(buildlist))
	if err != nil {
		return fmt.Errorf("parsing module list: %w", err)
	}

	missing := ""
	mods = gotool.FilterImportModules(mods)
	mods = gotool.RemoveExemptModules(mods, ExemptMods)
	prospects := adlr.MakeProspects(mods...)
	located, err := adlr.MakeProspector().Prospect(prospects...)
	if err != nil {
		missing = err.Error()
	}

	err = WriteJSONFile(locatedFile, located)
	if err != nil {
		return fmt.Errorf("writing located file: %w", err)
	}

	if missing != "" {
		return fmt.Errorf("%s", missing)
	}
	return nil
}
