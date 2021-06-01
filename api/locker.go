package api

import "github.com/blocky/adlr/internal"

type Locker = internal.Locker

func MakeLocker() Locker {
	return internal.MakeDependencyLocker()
}
