package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blocky/adlr"
	"github.com/spf13/cobra"
)

var IdentifiedFile string

var identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "Identify dependency license types",
	Long: `Outputs a file containing identified licenses and one
containing dependencies for which the license type could not be identified`,
	Run: func(cmd *cobra.Command, args []string) {
		err := Identify(LocatedFile, IdentifiedFile)
		ExitOnErr(err)
	},
}

func init() {
	identifyCmd.Flags().StringVarP(
		&LocatedFile, "located", "l", "./located-licenses.json",
		"Input file containing located licenses",
	)
	identifyCmd.Flags().StringVarP(
		&IdentifiedFile, "identified", "i", "./identified-licenses.json",
		"Output file containing identified licenses",
	)

	licenseCmd.AddCommand(identifyCmd)
}

func Identify(
	locatedFile string,
	identifiedFile string,
) error {
	locatedlist, err := os.Open(locatedFile)
	defer locatedlist.Close()
	if err != nil {
		return fmt.Errorf("opening located file: %w", err)
	}

	decoder := json.NewDecoder(locatedlist)
	var located []adlr.Mine
	err = decoder.Decode(&located)
	if err != nil {
		return fmt.Errorf("decoding located list: %w", err)
	}

	unidentified := ""
	identified, err := adlr.MakeMiner().Mine(located...)
	if err != nil {
		unidentified = err.Error()
	}

	bytes, err := json.MarshalIndent(identified, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling identified list: %w", err)
	}

	err = WriteFile(identifiedFile, bytes)
	if err != nil {
		return fmt.Errorf("writing identified file: %w", err)
	}

	if unidentified != "" {
		return fmt.Errorf("%s", unidentified)
	}
	return nil
}
