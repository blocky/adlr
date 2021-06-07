package api

import (
	"github.com/blocky/adlr/internal"
	"github.com/blocky/adlr/pkg/reader"
)

// Miner takes a variadic list of Mines and mines their license types
// and text, and returns an error containing all failures to deduce Mines'
// license type and text. The error is recoverable and purely for debugging,
// as LicenseLockManager will attempt to automatically solve them
type Miner interface {
	Mine(...Mine) ([]DependencyLock, error)
}

// MakeMiner creates a default Miner
func MakeMiner() Miner {
	return internal.MakeLicenseMiner()
}

// MakeMinerFromRaw creates a Miner with specified minimum confidences and
// a LimitedReader
func MakeMinerFromRaw(
	confidence float32,
	lead float32,
	reader *reader.LimitedReader,
) Miner {
	minimums := internal.Minimums{
		Confidence: confidence,
		Lead:       lead,
	}
	return internal.MakeLicenseMinerFromRaw(minimums, reader)
}
