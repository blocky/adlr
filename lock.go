package adlr

import (
	"bytes"
	"encoding/json"
)

type DependencyLock struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
	ErrStr  string  `json:"err,omitempty"`
	License License `json:"license"`
}

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

func (lock *DependencyLock) AddErrStr(
	errStr string,
) {
	lock.ErrStr = errStr
}

func DepLocksToDepLockMap(
	locks []DependencyLock,
) map[string]DependencyLock {
	var lockMap = make(map[string]DependencyLock, len(locks))
	for _, lock := range locks {
		lockMap[lock.Name] = lock
	}
	return lockMap
}

func MarshalDependencyLocks(
	locks []DependencyLock,
) ([]byte, error) {
	return json.Marshal(locks)
}

func UnmarshalDependencyLocks(
	bytes []byte,
) ([]DependencyLock, error) {
	var locks []DependencyLock
	err := json.Unmarshal(bytes, &locks)
	return locks, err
}

func DeserializeLocks(
	b []byte,
) ([]DependencyLock, error) {
	b = bytes.ReplaceAll(b, []byte("\\s"), []byte(" "))
	return UnmarshalDependencyLocks(b)
}

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
