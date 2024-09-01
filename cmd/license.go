package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "License commands",
	Long:  `Locate, identify, and verify dependency licenses`,
}

func init() {
	rootCmd.AddCommand(licenseCmd)
}

func ReadJSONFile(
	filename string,
	content any,
) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("reading bytes: %w", err)
	}

	err = json.Unmarshal(bytes, content)
	if err != nil {
		return fmt.Errorf("unmarshaling bytes: %w", err)
	}
	return nil
}

func WriteJSONFile(
	filename string,
	content any,
) error {
	bytes, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling bytes: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return fmt.Errorf("writing bytes: %w", err)
	}
	return nil
}
