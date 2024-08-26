package cmd

import (
	"fmt"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/pkg/ascertain"
	"github.com/blocky/adlr/pkg/reader"
	"github.com/spf13/cobra"
)

var IdentifiedFile string
var Confidence float32
var Lead float32

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
	identifyCmd.Flags().Float32VarP(
		&Confidence, "confidence", "c", ascertain.Confidence,
		"Minimum required confidence for a license match from text mining",
	)
	identifyCmd.Flags().Float32VarP(
		&Lead, "lead", "d", ascertain.Lead,
		"Minimum required difference between primary match and secondary matches")
	licenseCmd.AddCommand(identifyCmd)
}

func Identify(
	locatedFile string,
	identifiedFile string,
) error {
	var located []adlr.Mine
	err := ReadJSONFile(locatedFile, &located)
	if err != nil {
		return fmt.Errorf("reading located file: %w", err)
	}

	unidentified := ""
	lr := reader.NewLimitedReaderFromRaw(reader.Kilobyte * 1000)
	identified, err := adlr.MakeMinerFromRaw(Confidence, Lead, lr).Mine(located...)
	if err != nil {
		unidentified = err.Error()
	}

	err = WriteJSONFile(identifiedFile, identified)
	if err != nil {
		return fmt.Errorf("writing identified file: %w", err)
	}

	if unidentified != "" {
		return fmt.Errorf("%s", unidentified)
	}
	return nil
}
