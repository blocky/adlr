package adlr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

var newLockArray = []adlr.DependencyLock{
	adlr.DependencyLock{
		Name:    "1-happy-path",
		Version: "new-v1",
		License: adlr.License{
			Kind: "kind-1",
			Text: "text-1",
		},
	},
	adlr.DependencyLock{
		Name:    "2-missing-kind",
		Version: "new-v2",
		License: adlr.License{
			Kind: "",
			Text: "text-2",
		},
	},
	adlr.DependencyLock{
		Name:    "3-missing-text",
		Version: "new-v3",
		License: adlr.License{
			Kind: "kind-3",
			Text: "",
		},
	},
	adlr.DependencyLock{
		Name:    "4-missing-all",
		Version: "new-v4",
		License: adlr.License{
			Kind: "",
			Text: "",
		},
	},
}

var oldLockArray = []adlr.DependencyLock{
	adlr.DependencyLock{
		Name:    "1-happy-path",
		Version: "old-v1",
		License: adlr.License{
			Kind: "kind-1",
			Text: "text-1",
		},
	},
	adlr.DependencyLock{
		Name:    "2-missing-kind",
		Version: "old-v2",
		License: adlr.License{
			Kind: "kind-2",
			Text: "text-2",
		},
	},
	adlr.DependencyLock{
		Name:    "3-missing-text",
		Version: "old-v3",
		License: adlr.License{
			Kind: "kind-3",
			Text: "text-3",
		},
	},
	adlr.DependencyLock{
		Name:    "4-missing-all",
		Version: "old-v4",
		License: adlr.License{
			Kind: "kind-4",
			Text: "text-4",
		},
	},
}

func makeLockerErr(
	lock adlr.DependencyLock,
	errs ...error,
) adlr.LockerError {
	return adlr.MakeLockerError(lock.Name, lock.Version, errs...)
}

func TestDependencyLockerLockNew(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		locker := adlr.MakeDependencyLocker()
		locks := locker.LockNew(newLockArray)

		assert.Equal(t, newLockArray, locks)
	})
}

func TestDependencyLockerLockNewWithOld(t *testing.T) {
	newLocks := adlr.DepLocksToDepLockMap(newLockArray)
	oldLocks := adlr.DepLocksToDepLockMap(oldLockArray)

	locker := adlr.MakeDependencyLocker()
	result := locker.LockNewWithOld(
		newLocks,
		oldLocks,
	)

	resultsMap := adlr.DepLocksToDepLockMap(result)
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

	var locks = make([]adlr.DependencyLock, len(randomized))
	for i, name := range randomized {
		locks[i] = adlr.DependencyLock{Name: name}
	}

	locker := adlr.MakeDependencyLocker()
	result := locker.Alphabetize(locks)

	for j, lock := range result {
		assert.Equal(t, expected[j], lock.Name)
	}
}

func TestDependencyLockerVetLock(t *testing.T) {
	var kindErr = errors.New(adlr.ReqEditField + "kind")
	var textErr = errors.New(adlr.ReqEditField + "text")
	t.Run("happy path", func(t *testing.T) {
		lock := newLockArray[0]
		locker := adlr.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Nil(t, err)
	})
	t.Run("error on missing kind", func(t *testing.T) {
		lock := newLockArray[1]
		lockErr := makeLockerErr(lock, kindErr)
		locker := adlr.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
	t.Run("error on missing text", func(t *testing.T) {
		lock := newLockArray[2]
		lockErr := makeLockerErr(lock, textErr)
		locker := adlr.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
	t.Run("error on missing kind and text", func(t *testing.T) {
		lock := newLockArray[3]
		lockErr := makeLockerErr(lock, kindErr, textErr)
		locker := adlr.MakeDependencyLocker()
		err := locker.VetLock(lock)

		assert.Equal(t, lockErr, *err)
	})
}

func TestDependencyLockerVetLocks(t *testing.T) {
	var kindErr = errors.New(adlr.ReqEditField + "kind")
	t.Run("happy path", func(t *testing.T) {
		final := oldLockArray
		locker := adlr.MakeDependencyLocker()
		err := locker.VetLocks(final)
		assert.Nil(t, err)
	})
	t.Run("error on bad lock", func(t *testing.T) {
		badLock := newLockArray[1]
		final := append(oldLockArray, badLock)
		lockErr := makeLockerErr(badLock, kindErr)
		lockErrs := []adlr.LockerError{lockErr}

		locker := adlr.MakeDependencyLocker()
		err := locker.VetLocks(final)

		assert.Equal(t, lockErrs, err)
	})
}
