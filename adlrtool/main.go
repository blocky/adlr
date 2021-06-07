package main

import "github.com/blocky/adlr/adlrtool/cmd"

var DependencyRequirements = ""

func main() {
	cmd.DependencyRequirements = DependencyRequirements
	cmd.Execute()
}
