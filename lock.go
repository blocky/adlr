package adlr

import "encoding/json"

type DependencyLock struct {
	Name    string  `json:"name"`
	Version string  `json:"version"`
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

func DepLocksToDepLockMap(
	locks []DependencyLock,
) map[string]DependencyLock {
	var lockMap = make(map[string]DependencyLock, len(locks))
	for _, lock := range locks {
		lockMap[lock.Name] = lock
	}
	return lockMap
}

func UnmarshalDependencyLocks(
	bytes []byte,
) ([]DependencyLock, error) {
	var locks []DependencyLock
	err := json.Unmarshal(bytes, &locks)
	return locks, err
}
