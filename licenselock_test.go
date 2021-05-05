package adlr_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/internal/mocks"
	"github.com/blocky/adlr/reader"
	"github.com/blocky/prettyprinter"
)

var newLocksVettingErr = `[
 {
  "name": "bravo",
  "version": "v-bb",
  "errors": [
   "required editting of license field: kind"
  ]
 },
 {
  "name": "charlie",
  "version": "v-cb",
  "errors": [
   "required editting of license field: text"
  ]
 },
 {
  "name": "delta",
  "version": "v-db",
  "errors": [
   "required editting of license field: kind",
   "required editting of license field: text"
  ]
 }
]
`

func initLicenseLockMocks() (
	*mocks.Locker,
	string,
	*mocks.Printer,
	*reader.LimitedReader,
) {
	l := new(mocks.Locker)
	path := os.TempDir() + "/license.lock"
	p := new(mocks.Printer)
	r := reader.NewLimitedReader()
	return l, path, p, r
}

func TestLicenseLockExists(t *testing.T) {
	t.Run("false on lock file not exist", func(t *testing.T) {
		path := "./testdata/licenselock/unicorn.lock"
		locker, _, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)
		assert.False(t, lock.Exists())
	})
	t.Run("true on lock file exist", func(t *testing.T) {
		path := "./testdata/licenselock/old.lock"
		locker, _, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)
		assert.True(t, lock.Exists())
	})
}

func TestLicenseLockRead(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		path := "./testdata/licenselock/old.lock"
		locker, _, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)

		h := LicenseLockHelper{t, path}
		bytes := h.ReadFile("./testdata/licenselock/old.lock")
		expected := h.UnmarshalDependencyLocks(bytes)

		result, err := lock.Read()
		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})
	t.Run("error on bad lock json", func(t *testing.T) {
		locker, path, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)

		h := LicenseLockHelper{t, path}
		file := h.InitLock()
		defer h.CleanupLock()

		h.WriteFile(file, []byte(`{"bad":"data"}`))
		file.Close()
		jsonErr := "json: cannot unmarshal object into Go value of type []adlr.DependencyLock"

		_, err := lock.Read()
		assert.EqualError(t, err, jsonErr)
	})
}

func TestLicenseLockWriteAndVetLocks(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		locker, path, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)
		h := LicenseLockHelper{t, "./testdata/licenselock/old.lock"}
		bytes := h.ReadLock()
		locks := h.UnmarshalDependencyLocks(bytes)

		file := new(mocks.Writer)
		locker.
			On("VetLocks", locks).Return(nil)
		printer.
			On("Add", locks).Return(printer).
			On("Dump", file).Return(printer).
			On("Error").Return(nil)

		assert.Nil(t, lock.WriteAndVetLocks(file, locks))
		file.AssertExpectations(t)
		printer.AssertExpectations(t)
	})
	t.Run("error vetting locks", func(t *testing.T) {
		locker, path, printer, reader := initLicenseLockMocks()
		lock := adlr.MakeLicenseLockFromRaw(locker, path, printer, reader)
		h := LicenseLockHelper{t, "./testdata/licenselock/old.lock"}
		bytes := h.ReadLock()
		locks := h.UnmarshalDependencyLocks(bytes)

		file := new(mocks.Writer)
		printer.
			On("Add", locks).Return(printer).
			On("Dump", file).Return(printer).
			On("Error").Return(nil)

		lockErrs := make([]adlr.LockerError, 1)
		locker.
			On("VetLocks", locks).Return(lockErrs)
		printer.
			On("Add", lockErrs).Return(printer).
			On("StderrDump").Return(printer).
			On("Error").Return(nil)

		assert.Nil(t, lock.WriteAndVetLocks(file, locks))
		file.AssertExpectations(t)
		printer.AssertExpectations(t)
	})
}

func TestLicenseLockCreate(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		path := os.TempDir()
		lock := adlr.MakeLicenseLock(path)

		h := LicenseLockHelper{t, path + "/license.lock"}
		bytes := h.ReadFile("./testdata/licenselock/dependencies-old.json")
		deps := h.UnmarshalDependencyLocks(bytes)

		err := lock.Create(deps)
		defer h.CleanupLock()

		bytes = h.ReadLock()
		result := h.UnmarshalDependencyLocks(bytes)

		assert.Nil(t, err)
		assert.Equal(t, deps, result)
	})
	t.Run("errors on empty fields", func(t *testing.T) {
		lockPath := os.TempDir() + "/license.lock"
		stderrPath := os.TempDir() + "/stderr.tmp"
		h := LicenseLockHelper{t, lockPath}

		stderr := h.InitFile(stderrPath)
		defer h.CleanupFile(stderrPath)

		locker := adlr.MakeDependencyLocker()
		printer := prettyprinter.NewPrettyPrinterFromRaw(stderr, os.Stdout)
		reader := reader.NewLimitedReader()
		lock := adlr.MakeLicenseLockFromRaw(locker, lockPath, printer, reader)

		bytes := h.ReadFile("./testdata/licenselock/dependencies-new.json")
		deps := h.UnmarshalDependencyLocks(bytes)

		err := lock.Create(deps)
		defer h.CleanupLock()

		bytes = h.ReadLock()
		result := h.UnmarshalDependencyLocks(bytes)

		bytes = h.ReadFile(stderrPath)
		stderrErr := string(bytes)

		assert.Nil(t, err)
		assert.Equal(t, deps, result)
		assert.Equal(t, newLocksVettingErr, stderrErr)
	})
}

func TestLicenseLockOverwrite(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		path := os.TempDir()
		lock := adlr.MakeLicenseLock(path)

		h := LicenseLockHelper{t, path + "/license.lock"}
		file := h.InitLock()
		defer h.CleanupLock()

		expected := h.ReadFile("./testdata/licenselock/old.lock")
		h.WriteFile(file, expected)
		file.Close()

		bytes := h.ReadFile("./testdata/licenselock/dependencies-old.json")
		newDeps := h.UnmarshalDependencyLocks(bytes)
		err := lock.Overwrite(newDeps)
		assert.Nil(t, err)

		result := h.ReadLock()
		assert.Equal(t, string(expected), string(result))
	})
	t.Run("overwrite stress test", func(t *testing.T) {
		path := os.TempDir()
		lock := adlr.MakeLicenseLock(path)

		h := LicenseLockHelper{t, path + "/license.lock"}
		file := h.InitLock()
		defer h.CleanupLock()

		bytes := h.ReadFile("./testdata/licenselock/old.lock")
		h.WriteFile(file, bytes)
		file.Close()

		oldLocks := h.UnmarshalDependencyLocks(bytes)
		oldMap := adlr.DepLocksToDepLockMap(oldLocks)

		bytes = h.ReadFile("./testdata/licenselock/dependencies-new.json")
		newDeps := h.UnmarshalDependencyLocks(bytes)
		err := lock.Overwrite(newDeps)
		assert.Nil(t, err)

		bytes = h.ReadLock()
		resultLocks := h.UnmarshalDependencyLocks(bytes)
		resultMap := adlr.DepLocksToDepLockMap(resultLocks)

		alpha := resultMap["alpha"]
		bravo := resultMap["bravo"]
		charlie := resultMap["charlie"]
		delta := resultMap["delta"]

		oldBravo := oldMap["bravo"].License
		oldCharlie := oldMap["charlie"].License
		oldDelta := oldMap["delta"].License

		// check that new dependencies survive merge,
		// & empty new fields filled w/ old fields
		assert.Equal(t, "v-ab", alpha.Version)
		assert.Equal(t, "kind-a", alpha.License.Kind)
		assert.Equal(t, "text-a", alpha.License.Text)

		assert.Equal(t, "v-bb", bravo.Version)
		assert.Equal(t, oldBravo.Kind, bravo.License.Kind)
		assert.Equal(t, "text-b", bravo.License.Text)

		assert.Equal(t, "v-cb", charlie.Version)
		assert.Equal(t, "kind-c", charlie.License.Kind)
		assert.Equal(t, oldCharlie.Text, charlie.License.Text)

		assert.Equal(t, "v-db", delta.Version)
		assert.Equal(t, oldDelta.Kind, delta.License.Kind)
		assert.Equal(t, oldDelta.Text, delta.License.Text)
	})
	t.Run("overwriting but manual edits required", func(t *testing.T) {
		lockPath := os.TempDir() + "/license.lock"
		stderrPath := os.TempDir() + "/stderr.tmp"
		h := LicenseLockHelper{t, lockPath}

		stderr := h.InitFile(stderrPath)
		defer h.CleanupFile(stderrPath)

		locker := adlr.MakeDependencyLocker()
		printer := prettyprinter.NewPrettyPrinterFromRaw(stderr, os.Stdout)
		reader := reader.NewLimitedReader()
		lock := adlr.MakeLicenseLockFromRaw(locker, lockPath, printer, reader)

		file := h.InitLock()
		defer h.CleanupLock()

		expected := h.ReadFile("./testdata/licenselock/new.lock")
		h.WriteFile(file, expected)
		file.Close()

		bytes := h.ReadFile("./testdata/licenselock/dependencies-new.json")
		newDeps := h.UnmarshalDependencyLocks(bytes)
		err := lock.Overwrite(newDeps)
		assert.Nil(t, err)

		result := h.ReadLock()
		stderrErr := h.ReadFile(stderrPath)

		assert.Equal(t, newLocksVettingErr, string(stderrErr))
		assert.Equal(t, string(expected), string(result))
	})
}
