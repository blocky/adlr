package api

import "github.com/blocky/adlr/internal"

// Locker is a type alias for internal.Locker
type Locker = internal.Locker

// MakeLocker creates a Locker
func MakeLocker() Locker {
	return internal.MakeDependencyLocker()
}
