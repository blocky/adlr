package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/spf13/cobra"
)

var VerifiedFile string
var Whitelist []string

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify dependency licenses against whitelisted license types",
	Long:  `Outputs a file containing licenses verified against a whitelist`,
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
	identifiedList, err := os.Open(identifiedFile)
	defer identifiedList.Close()
	if err != nil {
		return fmt.Errorf("opening identified file: %w", err)
	}

	decoder := json.NewDecoder(identifiedList)
	var identified []adlr.DependencyLock
	err = decoder.Decode(&identified)
	if err != nil {
		return fmt.Errorf("decoding identified list: %w", err)
	}

	invalid := ""
	auditor := adlr.MakeAuditor(adlr.MakeWhitelist(Whitelist))
	verified, err := auditor.Audit(identified...)
	if err != nil {
		invalid = err.Error()
	}

	bytes, err := json.MarshalIndent(verified, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling verified list: %w", err)
	}

	err = WriteFile(verifiedFile, bytes)
	if err != nil {
		return fmt.Errorf("writing verified file: %w", err)
	}

	if invalid != "" {
		return fmt.Errorf("%s", invalid)
	}
	return nil
}
