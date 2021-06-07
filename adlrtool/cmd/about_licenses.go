package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/blocky/adlr"
	"github.com/blocky/prettyprinter"
)

var DependencyRequirements string
var DependencyNamesOnly bool
var DependencyName string

var aboutLicenseCmd = &cobra.Command{
	Use:   "license",
	Short: "A dependency license",
	Long:  `A software dependency license used by the client`,
	Run: func(cmd *cobra.Command, args []string) {
		locks, err := Deserialize(DependencyRequirements)
		ExitOnErr(err)

		PrintLicense(locks, DependencyName)
	},
}

var aboutLicensesCmd = &cobra.Command{
	Use:   "licenses",
	Short: "All dependency licenses",
	Long:  `List all software dependency licenses used by the client`,
	Run: func(cmd *cobra.Command, args []string) {
		locks, err := Deserialize(DependencyRequirements)
		ExitOnErr(err)

		if DependencyNamesOnly {
			PrintNames(locks)
			return
		}
		PrintLicenses(locks)
	},
}

func init() {
	nameKey := "name"
	namesKey := "names"

	aboutLicenseCmd.Flags().StringVarP(
		&DependencyName, nameKey, "n", "", "List a specific dependency license",
	)
	aboutLicensesCmd.Flags().BoolVar(
		&DependencyNamesOnly, namesKey, false, "List only dependency names",
	)

	aboutLicenseCmd.MarkFlagRequired(nameKey)

	aboutCmd.AddCommand(aboutLicenseCmd)
	aboutCmd.AddCommand(aboutLicensesCmd)
}

func Deserialize(
	deps string,
) ([]adlr.DependencyLock, error) {
	bytes := []byte(deps)
	return adlr.DeserializeLocks(bytes)
}

func PrintLicense(
	locks []adlr.DependencyLock,
	name string,
) {
	lockMap := adlr.DepLocksToDepLockMap(locks)
	lock, exist := lockMap[name]
	if !exist {
		ExitOnErr(
			errors.New("specified dependency does not exist: " + name),
		)
	}
	fmt.Fprintf(os.Stdout, "%s\n", lock.License.Text)
}

func PrintLicenses(
	locks []adlr.DependencyLock,
) {
	p := prettyprinter.NewPrettyPrinter()
	err := p.
		Add(locks).
		StdoutDump().
		StderrDumpOnError()
	ExitOnErr(err)
}

func PrintNames(
	locks []adlr.DependencyLock,
) {
	var names = make([]string, len(locks))
	for i, lock := range locks {
		names[i] = lock.Name
	}
	p := prettyprinter.NewPrettyPrinter()
	err := p.
		Add(names).
		StdoutDump().
		StderrDumpOnError()
	ExitOnErr(err)
}
