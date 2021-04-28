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

func TestDependencyToDependencyLock(t *testing.T) {
	m := gotool.Module{
		Path:    "/home/user/path",
		Version: "v1.0.1",
	}
	l := adlr.License{
		Kind: "MIT",
		Text: "license text",
	}
	d := adlr.Dependency{
		Module:  m,
		License: l,
	}
	dl := d.ToDependencyLock()

	assert.Equal(t, m.Path, dl.Name)
	assert.Equal(t, m.Version, dl.Version)
	assert.Equal(t, l, dl.License)
}
