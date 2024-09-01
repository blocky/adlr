package cmd

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var BuildlistFile string

var buildlistCmd = &cobra.Command{
	Use:   "buildlist",
	Short: "Build list of module dependencies",
	Long:  "Outputs a json file containing module dependencies.",
	Run: func(cmd *cobra.Command, args []string) {
		err := Build(BuildlistFile)
		ExitOnErr(err)
	},
}

func init() {
	buildlistCmd.Flags().StringVarP(
		&BuildlistFile, "buildlist", "b", "./buildlist.json",
		"Output file containing module dependencies",
	)
	licenseCmd.AddCommand(buildlistCmd)
}

func Build(
	buildlistFile string,
) error {
	listCmd := exec.Command("go", "list", "-m", "-json", "all")
	listOut, err := listCmd.Output()
	if err != nil {
		return fmt.Errorf("executing go list command: %w", err)
	}

	jqCmd := exec.Command("jq", "-s")
	jqCmd.Stdin = bytes.NewReader(listOut)
	jqOut, err := jqCmd.Output()
	if err != nil {
		return fmt.Errorf("executing jq command: %w", err)
	}

	err = WriteFile(buildlistFile, jqOut)
	if err != nil {
		return fmt.Errorf("writing buildlist file: %w", err)
	}
	return nil
}
