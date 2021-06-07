package api

import "github.com/blocky/adlr/pkg/ascertain"

// Locker is a type alias for ascertain.Locker
type Locker = ascertain.Locker

// MakeLocker creates a Locker
func MakeLocker() Locker {
	return ascertain.MakeDependencyLocker()
}
