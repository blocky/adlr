package adlr

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
