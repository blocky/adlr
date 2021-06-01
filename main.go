package main

import "github.com/blocky/adlr/cmd"

var DependencyRequirements = ""

func main() {
	cmd.DependencyRequirements = DependencyRequirements
	cmd.Execute()
}
