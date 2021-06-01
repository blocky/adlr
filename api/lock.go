package api

import "github.com/blocky/adlr/internal"

type DependencyLock = internal.DependencyLock

var DepLocksToDepLockMap = internal.DepLocksToDepLockMap

var MarshalDependencyLocks = internal.MarshalDependencyLocks

var DeserializeLocks = internal.DeserializeLocks

var SerializeLocks = internal.SerializeLocks
