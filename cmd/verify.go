package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	adlr "github.com/blocky/adlr/pkg"
)

var VerifiedFile string
var Whitelist []string

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify dependency licenses against whitelisted license types",
	Long:  "Outputs a file containing licenses verified against a whitelist",
	Run: func(cmd *cobra.Command, args []string) {
		err := Verify(IdentifiedFile, VerifiedFile)
		ExitOnErr(err)
	},
}

func init() {
	verifyCmd.Flags().StringVarP(
		&IdentifiedFile, "identified", "i", "./identified-licenses.json",
		"Input file containing identified licenses",
	)
	verifyCmd.Flags().StringVarP(
		&VerifiedFile, "verified", "v", "./verified-licenses.json",
		"Output file containing verified licenses",
	)
	verifyCmd.Flags().StringSliceVarP(
		&Whitelist, "whitelist", "w", adlr.DefaultWhitelist,
		"Comma separated list of acceptable licenses in SPDX License Identifier format",
	)
	licenseCmd.AddCommand(verifyCmd)
}

func Verify(
	identifiedFile string,
	verifiedFile string,
) error {
	var identified []adlr.DependencyLock
	err := ReadJSONFile(identifiedFile, &identified)
	if err != nil {
		return fmt.Errorf("reading identified file: %w", err)
	}

	invalid := ""
	auditor := adlr.MakeAuditor(adlr.MakeWhitelist(Whitelist))
	verified, err := auditor.Audit(identified...)
	if err != nil {
		invalid = err.Error()
	}

	err = WriteJSONFile(verifiedFile, verified)
	if err != nil {
		return fmt.Errorf("writing verified file: %w", err)
	}

	if invalid != "" {
		return fmt.Errorf("%s", invalid)
	}
	return nil
}
