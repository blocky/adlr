package main

import (
	_ "embed"

	"github.com/blocky/adlr/adlrtool/cmd"
)

//go:embed license.lock
var DependencyRequirements []byte

func main() {
	cmd.DependencyRequirements = DependencyRequirements
	cmd.Execute()
}
