package cmd

import (
	"fmt"
	"os"
	"os/exec"

	adlr "github.com/blocky/adlr/pkg"
	"github.com/blocky/adlr/pkg/gotool"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

var LocatedFile string
var ExemptMods []string

var locateCmd = &cobra.Command{
	Use:   "locate",
	Short: "Locate dependency licenses",
	Long:  "Outputs a json file containing located licenses.",
	Run: func(cmd *cobra.Command, args []string) {
		err := Locate(LocatedFile)
		ExitOnErr(err)
	},
}

func init() {
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
	locatedFile string,
) error {
	tidyCmd := exec.Command("go", "mod", "tidy")
	err := tidyCmd.Run()
	if err != nil {
		return fmt.Errorf("running go mod tidy: %w", err)
	}

	vendorCmd := exec.Command("go", "mod", "vendor")
	err = vendorCmd.Run()
	if err != nil {
		return fmt.Errorf("running go mod vendor: %w", err)
	}

	modFileBytes, err := os.ReadFile("./go.mod")
	if err != nil {
		return fmt.Errorf("reading go mod file: %w", err)
	}

	modFile, err := modfile.Parse("go.mod", modFileBytes, nil)
	if err != nil {
		return fmt.Errorf("parsing go mod file: %w", err)
	}

	// NOTE: This does not handle replaced or retracted modules
	mods := make([]gotool.Module, 0)
	for _, dep := range modFile.Require {
		mods = append(mods, gotool.Module{
			Path:     dep.Mod.Path,
			Dir:      "./vendor/" + dep.Mod.Path,
			Version:  dep.Mod.Version,
			Indirect: dep.Indirect,
		})
	}

	missing := ""
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

	// TODO: Should we remove the vendor folder?
	return nil
}
