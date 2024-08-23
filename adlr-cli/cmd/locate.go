package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/spf13/cobra"
)

var BuildlistFile string
var LocatedFile string

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate dependency licenses",
	Long: `Outputs a file containing located licenses. If there was trouble locating
one or more licenses, an error is returned with the list of missing licenses`,
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
	licenseCmd.AddCommand(locateCmd)
}

func Locate(
	buildlistFile string,
	locatedFile string,
) error {
	buildlist, err := os.Open(buildlistFile)
	defer buildlist.Close()
	if err != nil {
		return fmt.Errorf("opening buildlist file: %w", err)
	}

	parser := gotool.MakeBuildListParser()
	mods, err := parser.ParseModuleList(buildlist)
	if err != nil {
		return fmt.Errorf("parsing module list: %w", err)
	}

	missing := ""
	prospects := adlr.MakeProspects(gotool.FilterImportModules(mods)...)
	located, err := adlr.MakeProspector().Prospect(prospects...)
	if err != nil {
		missing = err.Error()
	}

	bytes, err := json.MarshalIndent(located, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling located list: %w", err)
	}

	err = WriteFile(locatedFile, bytes)
	if err != nil {
		return fmt.Errorf("writing located file: %w", err)
	}

	if missing != "" {
		return fmt.Errorf("%s", missing)
	}
	return nil
}
