package internal

import (
	"os"

	"github.com/blocky/adlr/reader"
	"github.com/blocky/prettyprinter"
)

// The filename for storing locked dependencies and their licenses
const LicenseLockName = "license.lock"

// LicenseLock is a file manager for locking dependencies and their
// licenses into an edittable file to save in a project's vcs
type LicenseLock struct {
	locker  Locker
	path    string
	printer prettyprinter.Printer
	reader  *reader.LimitedReader
}

// Create a LicenseLock with default values and providec directory
// to manage the lock file
func MakeLicenseLock(
	dir string,
) LicenseLock {
	locker := MakeDependencyLocker()
	path := dir + "/" + LicenseLockName
	printer := prettyprinter.NewPrettyPrinter()
	reader := reader.NewLimitedReaderFromRaw(reader.Kilobyte * 100)
	return LicenseLock{locker, path, printer, reader}
}

// Create a LicenseLock from specified values
func MakeLicenseLockFromRaw(
	locker Locker,
	path string,
	printer prettyprinter.Printer,
	reader *reader.LimitedReader,
) LicenseLock {
	return LicenseLock{
		locker:  locker,
		path:    path,
		printer: printer,
		reader:  reader,
	}
}

// Lock DependencyLocks derived by adlr of a project's golang
// module dependencies into a lock file for vcs
func (lock LicenseLock) Lock(
	locks []DependencyLock,
) error {
	if lock.Exists() {
		return lock.Overwrite(locks)
	}
	return lock.Create(locks)
}

// Create a lock file when one does not exist. Manual edits may be
// required
func (lock LicenseLock) Create(
	newLocks []DependencyLock,
) error {
	finalLocks := lock.locker.LockNew(newLocks)

	file, err := lock.OpenFileCreate()
	defer file.Close()
	if err != nil {
		return err
	}
	return lock.WriteAndVetLocks(file, finalLocks)
}

// Merge an existing lock file with new dependencies, giving new
// and updated dependencies priority. Manual edits may be required
func (lock LicenseLock) Overwrite(
	newLocks []DependencyLock,
) error {
	oldLocks, err := lock.Read()
	if err != nil {
		return err
	}
	file, err := lock.OpenFileOverwrite()
	defer file.Close()
	if err != nil {
		return err
	}
	finalLocks := lock.locker.
		LockNewWithOld(
			DepLocksToDepLockMap(newLocks),
			DepLocksToDepLockMap(oldLocks),
		)
	return lock.WriteAndVetLocks(file, finalLocks)
}

// Read and unmarshal an existing lock file into memory
func (lock LicenseLock) Read() (
	[]DependencyLock, error,
) {
	file, err := lock.OpenFileRead()
	defer file.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := lock.reader.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return UnmarshalDependencyLocks(bytes)
}

// Vet finalized merge locks for missing fields unsolvable
// by merging. Return an error for required manual edits
func (lock LicenseLock) VetLocks(
	locks []DependencyLock,
) error {
	lockErrs := lock.locker.VetLocks(locks)
	if len(lockErrs) != 0 {
		return lock.printer.
			Add(lockErrs).
			StderrDump().
			Error()
	}
	return nil
}

// Write locks to the lock file. Vet locks after writing for
// missing fields that require manual edits
func (lock LicenseLock) WriteAndVetLocks(
	writer prettyprinter.Writer,
	locks []DependencyLock,
) error {
	err := lock.Write(writer, locks)
	if err != nil {
		return err
	}
	// vet locks after write in case write fails
	return lock.VetLocks(locks)
}

// Write locks to lock file
func (lock LicenseLock) Write(
	writer prettyprinter.Writer,
	locks []DependencyLock,
) error {
	return lock.printer.
		Add(locks).
		Dump(writer).
		Error()

}

func (lock LicenseLock) OpenFileCreate() (*os.File, error) {
	var flag = os.O_RDWR | os.O_CREATE
	return os.OpenFile(lock.path, flag, 0666)
}

func (lock LicenseLock) OpenFileOverwrite() (*os.File, error) {
	var flag = os.O_RDWR | os.O_TRUNC
	return os.OpenFile(lock.path, flag, 0666)
}

func (lock LicenseLock) OpenFileRead() (*os.File, error) {
	return os.Open(lock.path)
}

func (lock LicenseLock) Exists() bool {
	_, pathErr := os.Stat(lock.path)
	return pathErr == nil
}
