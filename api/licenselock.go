package api

import (
	"github.com/blocky/adlr/internal"
	"github.com/blocky/adlr/reader"
	"github.com/blocky/prettyprinter"
)

type LicenseLockManager interface {
	Lock([]DependencyLock) error
	Read() ([]DependencyLock, error)
}

func MakeLicenseLockManager(
	dir string,
) LicenseLockManager {
	return internal.MakeLicenseLock(dir)
}

func MakeLicenseLockFromRaw(
	locker Locker,
	path string,
	printer prettyprinter.Printer,
	reader *reader.LimitedReader,
) LicenseLockManager {
	return internal.MakeLicenseLockFromRaw(
		locker,
		path,
		printer,
		reader,
	)
}
