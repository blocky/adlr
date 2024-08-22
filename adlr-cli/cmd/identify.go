package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/spf13/cobra"
)

var IdentifiedPath string

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "Identify dependency license types",
	Long: `Outputs a file containing identified licenses and one
containing dependencies for which the license type could not be identified`,
	Run: func(cmd *cobra.Command, args []string) {
		locatedlist, err := os.Open(LocatedPath)
		defer locatedlist.Close()
		ExitOnErr(err)
		Identify(locatedlist)
	},
}

func init() {
	identifyCmd.Flags().StringVarP(
		&LocatedPath, "located", "l", "./located-licenses.json",
		"Input file containing located licenses",
	)
	identifyCmd.Flags().StringVarP(
		&IdentifiedPath, "identified", "i", "./identified-licenses.json",
		"Output file containing identified licenses",
	)

	licenseCmd.AddCommand(identifyCmd)
}

func Identify(locatedlist *os.File) {
	decoder := json.NewDecoder(locatedlist)
	var located []adlr.Mine
	err := decoder.Decode(&located)
	if err != nil {
		fmt.Printf("decoding located: %w", err)
		os.Exit(1)
	}

	unidentified := ""
	identified, err := adlr.MakeMiner().Mine(located...)
	if err != nil {
		unidentified = err.Error()
	}

	bytes, err := json.MarshalIndent(identified, "", "\t")
	if err != nil {
		fmt.Printf("marshaling identified: %w", err)
		os.Exit(1)
	}

	err = WriteFile(IdentifiedPath, bytes)
	if err != nil {
		fmt.Printf("saving identified: %w", err)
		os.Exit(1)
	}

	if unidentified != "" {
		fmt.Printf("%s", unidentified)
		os.Exit(1)
	}
}
