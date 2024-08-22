package cmd

import (
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

func WriteFile(
	filename string,
	contents []byte,
) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(contents)
	if err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}
