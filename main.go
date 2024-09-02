package main

import (
	_ "embed"

	"github.com/blocky/adlr/cmd"
)

//go:embed version
var Version string

func main() {
	cmd.Version = Version
	cmd.Execute()
}
