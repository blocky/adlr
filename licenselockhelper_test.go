package adlr_test

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"testing"

	"github.com/blocky/adlr"
)

type LicenseLockHelper struct {
	t    *testing.T
	path string
}

func (l LicenseLockHelper) ReadLock() []byte {
	return l.ReadFile(l.path)
}

func (l LicenseLockHelper) ReadFile(
	path string,
) []byte {
	f, err := os.Open(path)
	l.checkError(err, "failed to open: "+path)

	bytes, err := ioutil.ReadAll(f)
	l.checkError(err, "failed to read: "+path)

	return bytes
}

func (l LicenseLockHelper) InitLock() *os.File {
	return l.InitFile(l.path)
}

func (l LicenseLockHelper) InitFile(
	path string,
) *os.File {
	file, err := os.Create(path)
	l.checkError(err, "failed to create: "+path)

	return file
}

func (l LicenseLockHelper) UnmarshalDependencyLocks(
	bytes []byte,
) []adlr.DependencyLock {
	locks, err := adlr.UnmarshalDependencyLocks(bytes)
	l.checkError(err, "failed to unmarshal locks")

	return locks
}

func (l LicenseLockHelper) UnmarshalDependencies(
	bytes []byte,
) []adlr.Dependency {
	var deps []adlr.Dependency
	err := json.Unmarshal(bytes, &deps)
	l.checkError(err, "failed to unmarshal dependencies")

	return deps
}

func (l LicenseLockHelper) WriteFile(
	file *os.File,
	bytes []byte,
) {
	_, err := file.Write(bytes)
	l.checkError(err, "failed to write: "+l.path)
}

func (l LicenseLockHelper) CleanupLock() {
	l.CleanupFile(l.path)
}

func (l LicenseLockHelper) CleanupFile(
	path string,
) {
	err := os.Remove(path)
	l.checkError(err, "failed to rm: "+path)
}

func (l LicenseLockHelper) checkError(
	err error,
	msg string,
) {
	if err != nil {
		l.t.Fatal(err, msg)
	}
}
