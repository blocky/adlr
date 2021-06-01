package api

import (
	"github.com/blocky/adlr/internal"
	"github.com/blocky/adlr/reader"
)

type Miner interface {
	Mine(...Mine) ([]DependencyLock, error)
}

func MakeMiner() Miner {
	return internal.MakeLicenseMiner()
}

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
