package internal_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/internal"
)

var newLockArray = []internal.DependencyLock{
	internal.DependencyLock{
		Name:    "1-happy-path",
		Version: "new-v1",
		License: internal.License{
			Kind: "kind-1",
			Text: "text-1",
		},
	},
	internal.DependencyLock{
		Name:    "2-missing-kind",
		Version: "new-v2",
		License: internal.License{
			Kind: "",
			Text: "text-2",
		},
	},
	internal.DependencyLock{
		Name:    "3-missing-text",
		Version: "new-v3",
		License: internal.License{
			Kind: "kind-3",
			Text: "",
		},
	},
	internal.DependencyLock{
		Name:    "4-missing-all",
		Version: "new-v4",
		License: internal.License{
			Kind: "",
			Text: "",
		},
	},
}

var oldLockArray = []internal.DependencyLock{
	internal.DependencyLock{
		Name:    "1-happy-path",
		Version: "old-v1",
		License: internal.License{
			Kind: "kind-1",
			Text: "text-1",
		},
	},
	internal.DependencyLock{
		Name:    "2-missing-kind",
		Version: "old-v2",
		License: internal.License{
			Kind: "kind-2",
			Text: "text-2",
		},
	},
	internal.DependencyLock{
		Name:    "3-missing-text",
		Version: "old-v3",
		License: internal.License{
			Kind: "kind-3",
			Text: "text-3",
		},
	},
	internal.DependencyLock{
		Name:    "4-missing-all",
		Version: "old-v4",
		License: internal.License{
			Kind: "kind-4",
			Text: "text-4",
		},
	},
}

func makeLockerErr(
	lock internal.DependencyLock,
	errs ...error,
) internal.LockerError {
	return internal.MakeLockerError(lock.Name, lock.Version, errs...)
}

func TestDependencyLockerLockNew(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		locker := internal.MakeDependencyLocker()
		locks := locker.LockNew(newLockArray)

		assert.Equal(t, newLockArray, locks)
	})
}

func TestDependencyLockerLockNewWithOld(t *testing.T) {
	newLocks := internal.DepLocksToDepLockMap(newLockArray)
	oldLocks := internal.DepLocksToDepLockMap(oldLockArray)

	locker := internal.MakeDependencyLocker()
	result := locker.LockNewWithOld(
		newLocks,
		oldLocks,
	)

	resultsMap := internal.DepLocksToDepLockMap(result)
	happyPath := resultsMap["1-happy-path"]
	missingKind := resultsMap["2-missing-kind"]
	missingText := resultsMap["3-missing-text"]
	missingAll := resultsMap["4-missing-all"]

	expected := newLockArray[0]
	assert.Equal(t, expected, happyPath)

	expected = oldLockArray[1]
	assert.Equal(t, expected.License.Kind, missingKind.License.Kind)
	assert.Equal(t, expected.License.Text, missingKind.License.Text)

	expected = oldLockArray[2]
	assert.Equal(t, expected.License.Kind, missingText.License.Kind)
	assert.Equal(t, expected.License.Text, missingText.License.Text)

	expected = oldLockArray[3]
	assert.Equal(t, expected.License.Kind, missingAll.License.Kind)
	assert.Equal(t, expected.License.Text, missingAll.License.Text)
}

func TestDependencyLockerAlphabetize(t *testing.T) {
	randomized := []string{"v", "u", "o", "q", "r", "y", "c", "x", "i", "z", "g", "j",
		"f", "p", "w", "b", "m", "k", "e", "n", "l", "d", "t", "a", "s", "h"}
	expected := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var locks = make([]internal.DependencyLock, len(randomized))
	for i, name := range randomized {
		locks[i] = internal.DependencyLock{Name: name}
	}

	locker := internal.MakeDependencyLocker()
	result := locker.Alphabetize(locks)

	for j, lock := range result {
		assert.Equal(t, expected[j], lock.Name)
	}
}

func TestDependencyLockerVetLock(t *testing.T) {
	var kindErr = errors.New(internal.ReqEditField + "kind")
	var textErr = errors.New(internal.ReqEditField + "text")
	t.Run("happy path", func(t *testing.T) {
		lock := newLockArray[0]
		locker := internal.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Nil(t, err)
	})
	t.Run("error on missing kind", func(t *testing.T) {
		lock := newLockArray[1]
		lockErr := makeLockerErr(lock, kindErr)
		locker := internal.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
	t.Run("error on missing text", func(t *testing.T) {
		lock := newLockArray[2]
		lockErr := makeLockerErr(lock, textErr)
		locker := internal.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
	t.Run("error on missing kind and text", func(t *testing.T) {
		lock := newLockArray[3]
		lockErr := makeLockerErr(lock, kindErr, textErr)
		locker := internal.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
}

func TestDependencyLockerVetLocks(t *testing.T) {
	var kindErr = errors.New(internal.ReqEditField + "kind")
	t.Run("happy path", func(t *testing.T) {
		final := oldLockArray
		locker := internal.MakeDependencyLocker()
		err := locker.VetLocks(final)
		assert.Nil(t, err)
	})
	t.Run("error on bad lock", func(t *testing.T) {
		badLock := newLockArray[1]
		final := append(oldLockArray, badLock)
		lockErr := makeLockerErr(badLock, kindErr)
		lockErrs := []internal.LockerError{lockErr}

		locker := internal.MakeDependencyLocker()
		err := locker.VetLocks(final)

		assert.Equal(t, lockErrs, err)
	})
}
