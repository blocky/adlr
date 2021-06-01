package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/internal"
	"github.com/blocky/adlr/reader"
)

var DependencyLocks = []internal.DependencyLock{
	internal.DependencyLock{
		Name:    "github.com/spf13/viper",
		Version: "v1.4.0",
		License: internal.License{
			"MIT",
			"MIT License",
		},
	},
	internal.DependencyLock{
		Name:    "github.com/stretchr/testify",
		Version: "v1.6.1",
		License: internal.License{
			"MIT",
			"MIT License",
		},
	},
	internal.DependencyLock{
		Name:    "github.com/ybbus/jsonrpc",
		Version: "v2.1.2+incompatible",
		License: internal.License{
			"MIT",
			"MIT License",
		},
	},
}

func TestMakeDependencyLock(t *testing.T) {
	l := internal.MakeLicense("kind", "text")
	dl := internal.MakeDependencyLock("name", "version", l)

	assert.Equal(t, l, dl.License)
	assert.Equal(t, "name", dl.Name)
	assert.Equal(t, "version", dl.Version)
}

func TestDependencyLockAddErrStr(t *testing.T) {
	var l internal.DependencyLock
	l.AddErrStr("error")

	assert.Equal(t, "error", l.ErrStr)
}

func TestDepLocksToDepLocksMap(t *testing.T) {
	dl1 := internal.MakeDependencyLock("n1", "v1", internal.MakeLicense("k1", "t1"))
	dl2 := internal.MakeDependencyLock("n2", "v2", internal.MakeLicense("k2", "t2"))
	dl3 := internal.MakeDependencyLock("n3", "v3", internal.MakeLicense("k3", "t3"))

	dlSlice := []internal.DependencyLock{dl1, dl2, dl3}
	dlMap := internal.DepLocksToDepLockMap(dlSlice)

	assert.Equal(t, dl1, dlMap[dl1.Name])
	assert.Equal(t, dl2, dlMap[dl2.Name])
	assert.Equal(t, dl3, dlMap[dl3.Name])
}

func TestMarshalDependencyLocks(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		bytes, err := reader.
			NewLimitedReader().
			ReadFileFromPath("./testdata/lock/marshaled.json")
		assert.Nil(t, err)
		result, err := internal.MarshalDependencyLocks(DependencyLocks)

		assert.Nil(t, err)
		assert.Equal(t, string(bytes), string(result))
	})
}

func TestUnmarshalDependencyLocks(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		bytes, err := reader.
			NewLimitedReader().
			ReadFileFromPath("./testdata/lock/marshaled.json")
		assert.Nil(t, err)
		result, err := internal.UnmarshalDependencyLocks(bytes)

		assert.Nil(t, err)
		assert.Equal(t, DependencyLocks, result)
	})
	t.Run("error on unmarshaling", func(t *testing.T) {
		bytes := []byte("{\"bad\":\"json\"}")
		_, err := internal.UnmarshalDependencyLocks(bytes)

		assert.EqualError(t, err,
			"json: cannot unmarshal object into "+
				"Go value of type []internal.DependencyLock")
	})
}

func TestLocksSerialization(t *testing.T) {
	reader := reader.NewLimitedReader()

	bytes, err := reader.ReadFileFromPath("./testdata/lock/deserialized.json")
	assert.Nil(t, err)
	expected, err := internal.UnmarshalDependencyLocks(bytes)
	assert.Nil(t, err)

	bytes, err = reader.ReadFileFromPath("./testdata/lock/serialized.txt")
	assert.Nil(t, err)
	result, err := internal.DeserializeLocks(bytes)
	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}
