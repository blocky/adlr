package api

import (
	"github.com/blocky/adlr/internal"
	"github.com/blocky/adlr/reader"
	"github.com/blocky/prettyprinter"
)

// LicenseLockManager is a file manager for the license.lock file.
// Lock() will create a new file or merge with an existing.
// Read() will return a list of DependencyLocks, if a file exists.
type LicenseLockManager interface {
	Lock(...DependencyLock) error
	Read() ([]DependencyLock, error)
}

// MakeLicenseLockManager creates a LicenseLockManager with specified directory
func MakeLicenseLockManager(
	dir string,
) LicenseLockManager {
	return internal.MakeLicenseLock(dir)
}

// MakeLicenseLockFromRaw creates a LicenseLockManager from specified parameters
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