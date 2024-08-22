package main

import (
	_ "embed"

	"github.com/blocky/adlr/adlr-cli/cmd"
)

//go:embed version
var Version string

//go:embed license.lock
var DependencyRequirements []byte

func main() {
	cmd.Version = Version
	cmd.DependencyRequirements = DependencyRequirements
	cmd.Execute()
}
