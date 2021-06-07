package internal

import "github.com/blocky/prettyprinter"

// For mockery to create a mock for testing
type Printer interface {
	prettyprinter.Printer
}

// For mockery to create a mock for testing
type Writer interface {
	prettyprinter.Writer
}
