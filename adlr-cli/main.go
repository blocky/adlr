package main

import (
	_ "embed"

	"github.com/blocky/adlr/adlr-cli/cmd"
)

//go:embed version
var Version string

func main() {
	cmd.Version = Version
	cmd.Execute()
}
