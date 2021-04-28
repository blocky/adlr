package adlr

import "errors"

const (
	ReqEditField = "required editting of license field: "
)

type Locker interface {
	LockNew([]DependencyLock) []DependencyLock
	LockNewWithOld(new, old map[string]DependencyLock) []DependencyLock
	VetLocks([]DependencyLock) []LockerError
}

type DependencyLocker struct{}

func MakeDependencyLocker() DependencyLocker {
	return DependencyLocker{}
}

func (l DependencyLocker) LockNew(
	new []DependencyLock,
) []DependencyLock {
	return l.Alphabetize(new)
}

func (l DependencyLocker) LockNewWithOld(
	new, old map[string]DependencyLock,
) []DependencyLock {
	var final = make([]DependencyLock, len(new))
	var i int
	for name, newlock := range new {
		oldlock, exists := old[name]
		if exists {
			// match, merge
			if newlock.License.Kind == "" {
				newlock.License.Kind = oldlock.License.Kind
			}
			if newlock.License.Text == "" {
				newlock.License.Text = oldlock.License.Text
			}
		}
		final[i] = newlock
		i++
		continue
	}
	return l.Alphabetize(final)
}

func (l DependencyLocker) Alphabetize(
	locks []DependencyLock,
) []DependencyLock {
	var length = len(locks)

	if length == 1 {
		return locks
	}

	middle := int(length / 2)
	left := make([]DependencyLock, middle)
	right := make([]DependencyLock, length-middle)

	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = locks[i]
		} else {
			right[i-middle] = locks[i]
		}
	}
	return merge(
		l.Alphabetize(left),
		l.Alphabetize(right),
	)
}

func merge(
	left, right []DependencyLock,
) []DependencyLock {
	var merged = make([]DependencyLock, len(left)+len(right))

	var i int
	for len(left) > 0 && len(right) > 0 {
		if left[0].Name < right[0].Name {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		merged[i] = left[j]
		i++
	}
	for k := 0; k < len(right); k++ {
		merged[i] = right[k]
		i++
	}
	return merged
}

func (l DependencyLocker) VetLocks(
	final []DependencyLock,
) []LockerError {
	// check finalized locks for fields requiring edits
	var lockErrs []LockerError
	for _, lock := range final {

		lockErr := l.VetLock(lock)
		if lockErr != nil {
			lockErrs = append(lockErrs, *lockErr)
		}
	}
	if len(lockErrs) != 0 {
		return lockErrs
	}
	return nil
}

func (l DependencyLocker) VetLock(
	lock DependencyLock,
) *LockerError {
	var errs []error
	if lock.License.Kind == "" {
		errs = append(errs, errors.New(ReqEditField+"kind"))
	}
	if lock.License.Text == "" {
		errs = append(errs, errors.New(ReqEditField+"text"))
	}
	if len(errs) != 0 {
		lockErr := MakeLockerError(
			lock.Name,
			lock.Version,
			errs...,
		)
		return &lockErr
	}
	return nil
}
