package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/reader"
)

var DependencyLocks = []adlr.DependencyLock{
	adlr.DependencyLock{
		"github.com/spf13/viper",
		"v1.4.0",
		adlr.License{
			"MIT",
			"MIT License",
		},
	},
	adlr.DependencyLock{
		"github.com/stretchr/testify",
		"v1.6.1",
		adlr.License{
			"MIT",
			"MIT License",
		},
	},
	adlr.DependencyLock{
		"github.com/ybbus/jsonrpc",
		"v2.1.2+incompatible",
		adlr.License{
			"MIT",
			"MIT License",
		},
	},
}

func TestMakeDependencyLock(t *testing.T) {
	l := adlr.MakeLicense("kind", "text")
	dl := adlr.MakeDependencyLock("name", "version", l)

	assert.Equal(t, l, dl.License)
	assert.Equal(t, "name", dl.Name)
	assert.Equal(t, "version", dl.Version)
}

func TestDepLocksToDepLocksMap(t *testing.T) {
	dl1 := adlr.MakeDependencyLock("n1", "v1", adlr.MakeLicense("k1", "t1"))
	dl2 := adlr.MakeDependencyLock("n2", "v2", adlr.MakeLicense("k2", "t2"))
	dl3 := adlr.MakeDependencyLock("n3", "v3", adlr.MakeLicense("k3", "t3"))

	dlSlice := []adlr.DependencyLock{dl1, dl2, dl3}
	dlMap := adlr.DepLocksToDepLockMap(dlSlice)

	assert.Equal(t, dl1, dlMap[dl1.Name])
	assert.Equal(t, dl2, dlMap[dl2.Name])
	assert.Equal(t, dl3, dlMap[dl3.Name])
}

func TestUnmarshalDependencyLocks(t *testing.T) {
	bytes, err := reader.
		NewLimitedReader().
		ReadFileFromPath("./testdata/lock/deserialized.txt")
	assert.Nil(t, err)
	result, err := adlr.UnmarshalDependencyLocks(bytes)
	assert.Nil(t, err)

	assert.Equal(t, DependencyLocks, result)
}
