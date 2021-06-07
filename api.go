package adlr

import (
	"github.com/blocky/adlr/pkg/ascertain"
	"github.com/blocky/adlr/pkg/reader"
	"github.com/blocky/prettyprinter"
)

var (
	// DepLocksToDepLockMap is function alias for ascertain.DepLocksToDepLockMap
	DepLocksToDepLockMap = ascertain.DepLocksToDepLockMap

	// MarshalDependencyLocks is function alias for ascertain.MarshalDependencyLocks
	MarshalDependencyLocks = ascertain.MarshalDependencyLocks

	// MakeProspects is a type alias for ascertain.MakeProspects
	MakeProspects = ascertain.MakeProspects

	// UnmarshalDependencyLocks is a function alias for ascertain.UnmarshalDependencyLocks
	UnmarshalDependencyLocks = ascertain.UnmarshalDependencyLocks

	// DeserializeLocks is function alias for ascertain.DeserializeLocks
	DeserializeLocks = ascertain.DeserializeLocks

	// SerializeLocks is function alias for ascertain.SerializeLocks
	SerializeLocks = ascertain.SerializeLocks
)

// Prospect is a type alias for ascertain.Prospect
type Prospect = ascertain.Prospect

// Prospector takes a variadic list of Prospects and uses text mining to
// derive Mines with potential matches containing license type, license
// file name, and confidence float, and returns an error of failed Mines
// from missing directories or missing license files
type Prospector interface {
	Prospect(...Prospect) ([]Mine, error)
}

// MakeProspector creates a Prospector
func MakeProspector() Prospector {
	return ascertain.MakeLicenseProspector()
}

// DependencyLock is a type alias for ascertain.DependencyLock
type DependencyLock = ascertain.DependencyLock

// Mine is a type alias for ascertain.Mine
type Mine = ascertain.Mine

// Miner takes a variadic list of Mines and mines their license types
// and text, and returns an error containing all failures to deduce Mines'
// license type and text. The error is recoverable and purely for debugging,
// as LicenseLockManager will attempt to automatically solve them
type Miner interface {
	Mine(...Mine) ([]DependencyLock, error)
}

// MakeMiner creates a default Miner
func MakeMiner() Miner {
	return ascertain.MakeLicenseMiner()
}

// MakeMinerFromRaw creates a Miner with specified minimum confidences and
// a LimitedReader
func MakeMinerFromRaw(
	confidence float32,
	lead float32,
	reader *reader.LimitedReader,
) Miner {
	minimums := ascertain.Minimums{
		Confidence: confidence,
		Lead:       lead,
	}
	return ascertain.MakeLicenseMinerFromRaw(minimums, reader)
}

// Locker is a type alias for ascertain.Locker
type Locker = ascertain.Locker

// MakeLocker creates a Locker
func MakeLocker() Locker {
	return ascertain.MakeDependencyLocker()
}

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
	return ascertain.MakeLicenseLock(dir)
}

// MakeLicenseLockFromRaw creates a LicenseLockManager from specified parameters
func MakeLicenseLockFromRaw(
	locker Locker,
	path string,
	printer prettyprinter.Printer,
	reader *reader.LimitedReader,
) LicenseLockManager {
	return ascertain.MakeLicenseLockFromRaw(
		locker,
		path,
		printer,
		reader,
	)
}

// This list is of SPDX License Identifiers, the standard
// used by the text-mining package:
// github.com/go-enry/go-license-detector/v4
// and contains licenses deemed automatically fulfillable.
// To add to this list, see:
// https://spdx.org/licenses/
// for license identifiers
var DefaultWhitelist = []string{
	"Apache-2.0",
	"BSD-1-Clause",
	"BSD-2-Clause",
	"BSD-3-Clause",
	"MIT",
	"MIT-0",
}

// Whitelist is a type alias for ascertain.Whitelist
type Whitelist = ascertain.Whitelist

// MakeWhitelist creates a Whitelist from specified licenses
func MakeWhitelist(
	licenses []string,
) Whitelist {
	return ascertain.MakeLicenseWhitelist(licenses)
}

// Auditor takes a variadic list of DependencyLocks and audits their licenses
// against a license whitelist, returning an error of all offending
// DependencyLocks and their non-whitelisted licenses
type Auditor interface {
	Audit(...DependencyLock) error
}

// MakeAuditor creates an Auditor with a specified Whitelist
func MakeAuditor(
	whitelist Whitelist,
) Auditor {
	return ascertain.MakeLicenseAuditor(whitelist)
}
