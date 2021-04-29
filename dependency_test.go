package adlr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/gotool"
)

func TestMakeDependency(t *testing.T) {
	m1 := gotool.Module{Path: "/home/user/path1"}
	m2 := gotool.Module{Path: "/home/user/path2"}
	m3 := gotool.Module{Path: "/home/user/path3"}
	d := adlr.MakeDependencies(m1, m2, m3)

	assert.True(t, len(d) == 3)
	assert.Equal(t, d[0].Module, m1)
	assert.Equal(t, d[1].Module, m2)
	assert.Equal(t, d[2].Module, m3)
}

func TestDependencyAddError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		var d adlr.Dependency
		err := errors.New("error")
		d.AddError(err)

		assert.Equal(t, "error", d.ErrStr)
	})
	t.Run("nil error", func(t *testing.T) {
		var d adlr.Dependency
		var err error = nil
		d.AddError(err)

		assert.Equal(t, "", d.ErrStr)
	})
}

func TestDepsToDepLockArray(t *testing.T) {
	d1 := adlr.Dependency{
		Module:  gotool.Module{Path: "p1", Version: "v1"},
		License: adlr.MakeLicense("k1", "t1"),
	}
	d2 := adlr.Dependency{
		Module:  gotool.Module{Path: "p2", Version: "v2"},
		License: adlr.MakeLicense("k2", "t2"),
	}
	d3 := adlr.Dependency{
		Module:  gotool.Module{Path: "p3", Version: "v3"},
		License: adlr.MakeLicense("k3", "t3"),
	}
	dSlice := []adlr.Dependency{d1, d2, d3}
	dlSlice := adlr.DepsToDepLockArray(dSlice)

	for i, d := range dSlice {
		var dl = dlSlice[i]
		assert.Equal(t, d.Module.Path, dl.Name)
		assert.Equal(t, d.Module.Version, dl.Version)
		assert.Equal(t, d.License, dl.License)
	}
}

func TestDepsToDepLockMap(t *testing.T) {
	d1 := adlr.Dependency{
		Module:  gotool.Module{Path: "p1", Version: "v1"},
		License: adlr.MakeLicense("k1", "t1"),
	}
	d2 := adlr.Dependency{
		Module:  gotool.Module{Path: "p2", Version: "v2"},
		License: adlr.MakeLicense("k2", "t2"),
	}
	d3 := adlr.Dependency{
		Module:  gotool.Module{Path: "p3", Version: "v3"},
		License: adlr.MakeLicense("k3", "t3"),
	}
	dSlice := []adlr.Dependency{d1, d2, d3}
	dlMap := adlr.DepsToDepLockMap(dSlice)

	for _, d := range dSlice {
		var dl = dlMap[d.Module.Path]
		assert.Equal(t, d.Module.Path, dl.Name)
		assert.Equal(t, d.Module.Version, dl.Version)
		assert.Equal(t, d.License, dl.License)
	}
}
