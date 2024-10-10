package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"

	adlr "github.com/blocky/adlr/pkg"
	"github.com/blocky/adlr/pkg/gotool"
)

var LocatedFile string
var ProjectDir string
var ExemptMods []string

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate dependency licenses",
	Long:  "Outputs a json file containing located licenses.",
	Run: func(cmd *cobra.Command, args []string) {
		err := Locate(LocatedFile, ProjectDir)
		ExitOnErr(err)
	},
}

func init() {
	locateCmd.Flags().StringVarP(
		&LocatedFile, "located", "l", "./located-licenses.json",
		"Output file containing located licenses",
	)
	locateCmd.Flags().StringVarP(
		&ProjectDir, "project-dir", "p", "./",
		"Path to the Go project whose dependencies you wish to audit",
	)
	locateCmd.Flags().StringSliceVarP(
		&ExemptMods, "exempt-modules", "e", []string{},
		"Comma separated list of modules to ignore during license location")
	licenseCmd.AddCommand(locateCmd)
}

func Locate(
	locatedFile string,
	projectDir string,
) error {
	modFileBytes, err := os.ReadFile(projectDir + "/go.mod")
	if err != nil {
		return fmt.Errorf("reading go mod file: %w", err)
	}

	modFile, err := modfile.Parse("go.mod", modFileBytes, nil)
	if err != nil {
		return fmt.Errorf("parsing go mod file: %w", err)
	}

	replace := make(map[string]module.Version, 0)
	for _, mod := range modFile.Replace {
		replace[mod.Old.Path] = mod.New
	}

	require := make([]gotool.Module, 0)
	for _, mod := range modFile.Require {
		if val, ok := replace[mod.Mod.Path]; ok {
			mod.Mod.Path = val.Path
			mod.Mod.Version = val.Version

		}

		require = append(require, gotool.Module{
			Path:     mod.Mod.Path,
			Dir:      projectDir + "/vendor/" + mod.Mod.Path,
			Version:  mod.Mod.Version,
			Indirect: mod.Indirect,
		})
	}

	missing := ""
	require = gotool.RemoveExemptModules(require, ExemptMods)
	prospects := adlr.MakeProspects(require...)
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
