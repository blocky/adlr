package adlr

import (
	"github.com/go-enry/go-license-detector/v4/licensedb"

	"github.com/blocky/adlr/gotool"
)

type Dependency struct {
	Module  gotool.Module
	Result  licensedb.Result
	License License
	ErrStr  string
}

func MakeDependencies(
	modules ...gotool.Module,
) []Dependency {
	var deps = make([]Dependency, len(modules))
	for i, m := range modules {
		deps[i].Module = m
	}
	return deps
}

func (d *Dependency) AddResult(
	result licensedb.Result,
) {
	d.Result = result
}

func (d *Dependency) AddLicense(
	license License,
) {
	d.License = license
}

func (d *Dependency) AddError(
	err error,
) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	d.ErrStr = errStr
}

func (d Dependency) ToDependencyLock() DependencyLock {
	return MakeDependencyLock(
		d.Module.Path,
		d.Module.Version,
		d.License,
	)
}

func DepsToDepLockArray(
	deps []Dependency,
) []DependencyLock {
	var locks = make([]DependencyLock, len(deps))
	for i, dep := range deps {
		locks[i] = dep.ToDependencyLock()
	}
	return locks
}

func DepsToDepLockMap(
	deps []Dependency,
) map[string]DependencyLock {
	var lockMap = make(map[string]DependencyLock, len(deps))
	for _, dep := range deps {
		lock := dep.ToDependencyLock()
		lockMap[lock.Name] = lock
	}
	return lockMap
}
