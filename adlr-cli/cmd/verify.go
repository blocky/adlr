package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/spf13/cobra"
)

var VerifiedPath string
var Whitelist []string

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify dependency licenses against whitelisted license types",
	Long:  `Outputs a file containing licenses verified against a whitelist`,
	Run: func(cmd *cobra.Command, args []string) {
		identifiedList, err := os.Open(IdentifiedPath)
		defer identifiedList.Close()
		ExitOnErr(err)
		Verify(identifiedList)
	},
}

func init() {
	verifyCmd.Flags().StringVarP(
		&IdentifiedPath, "identified", "i", "./identified-licenses.json",
		"Input file containing identified licenses",
	)
	verifyCmd.Flags().StringVarP(
		&VerifiedPath, "verified", "v", "./verified-licenses.json",
		"Output file containing verified licenses",
	)
	verifyCmd.Flags().StringSliceVarP(
		&Whitelist, "whitelist", "w", adlr.DefaultWhitelist,
		"Comma separated list of acceptable licenses in SPDX License Identifier format",
	)
	licenseCmd.AddCommand(verifyCmd)
}

func Verify(identifiedList *os.File) {
	decoder := json.NewDecoder(identifiedList)
	var identified []adlr.DependencyLock
	err := decoder.Decode(&identified)
	if err != nil {
		fmt.Printf("decoding identified: %w", err)
		os.Exit(1)
	}

	invalid := ""
	auditor := adlr.MakeAuditor(adlr.MakeWhitelist(Whitelist))
	verified, err := auditor.Audit(identified...)
	if err != nil {
		invalid = err.Error()
	}

	bytes, err := json.MarshalIndent(verified, "", "\t")
	if err != nil {
		fmt.Printf("marshaling verified: %w", err)
		os.Exit(1)
	}

	err = WriteFile(VerifiedPath, bytes)
	if err != nil {
		fmt.Printf("saving verified: %w", err)
		os.Exit(1)
	}

	if invalid != "" {
		fmt.Printf("%s", invalid)
		os.Exit(1)
	}
}
