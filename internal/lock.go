package internal

import (
	"bytes"
	"encoding/json"
)

// DependencyLock can be written directly to the license.lock file
// when finalized
type DependencyLock struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	ErrStr  string  `json:"err,omitempty"`
	License License `json:"license"`
}

// Create a DependencyLock
func MakeDependencyLock(
	name, version string,
	license License,
) DependencyLock {
	return DependencyLock{
		Name:    name,
		Version: version,
		License: license,
	}
}

// Add error string
func (lock *DependencyLock) AddErrStr(
	errStr string,
) {
	lock.ErrStr = errStr
}

// Mutate a DependencyLock array to a DependencyLock map
func DepLocksToDepLockMap(
	locks []DependencyLock,
) map[string]DependencyLock {
	var lockMap = make(map[string]DependencyLock, len(locks))
	for _, lock := range locks {
		lockMap[lock.Name] = lock
	}
	return lockMap
}

// Marshal a list of DependencyLock
func MarshalDependencyLocks(
	locks []DependencyLock,
) ([]byte, error) {
	return json.Marshal(locks)
}

// Unmarshal a list of DependencyLock
func UnmarshalDependencyLocks(
	bytes []byte,
) ([]DependencyLock, error) {
	var locks []DependencyLock
	err := json.Unmarshal(bytes, &locks)
	return locks, err
}

// Deserialzed a list of DependencyLocks after inclusion into
// a variable from building with golang -ldflags
func DeserializeLocks(
	b []byte,
) ([]DependencyLock, error) {
	b = bytes.ReplaceAll(b, []byte("\\s"), []byte(" "))
	return UnmarshalDependencyLocks(b)
}

// Serialize a list of DependencyLocks for inclusion into a
// variable from golang -ldflags. Golang buildflags require
// no newlines or spaces
func SerializeLocks(
	locks []DependencyLock,
) ([]byte, error) {
	b, err := MarshalDependencyLocks(locks)
	if err != nil {
		return nil, err
	}
	// go ldflags require no spaces or newline chars
	b = bytes.ReplaceAll(b, []byte("\n"), []byte(""))
	b = bytes.ReplaceAll(b, []byte(" "), []byte("\\s"))
	return b, nil
}
