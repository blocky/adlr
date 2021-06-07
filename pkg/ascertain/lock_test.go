package ascertain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
	"github.com/blocky/adlr/pkg/reader"
)

var DependencyLocks = []ascertain.DependencyLock{
	ascertain.DependencyLock{
		Name:    "github.com/spf13/viper",
		Version: "v1.4.0",
		License: ascertain.License{
			"MIT",
			"MIT License",
		},
	},
	ascertain.DependencyLock{
		Name:    "github.com/stretchr/testify",
		Version: "v1.6.1",
		License: ascertain.License{
			"MIT",
			"MIT License",
		},
	},
	ascertain.DependencyLock{
		Name:    "github.com/ybbus/jsonrpc",
		Version: "v2.1.2+incompatible",
		License: ascertain.License{
			"MIT",
			"MIT License",
		},
	},
}

func TestMakeDependencyLock(t *testing.T) {
	l := ascertain.MakeLicense("kind", "text")
	dl := ascertain.MakeDependencyLock("name", "version", l)

	assert.Equal(t, l, dl.License)
	assert.Equal(t, "name", dl.Name)
	assert.Equal(t, "version", dl.Version)
}

func TestDependencyLockAddErrStr(t *testing.T) {
	var l ascertain.DependencyLock
	l.AddErrStr("error")

	assert.Equal(t, "error", l.ErrStr)
}

func TestDepLocksToDepLocksMap(t *testing.T) {
	dl1 := ascertain.MakeDependencyLock("n1", "v1", ascertain.MakeLicense("k1", "t1"))
	dl2 := ascertain.MakeDependencyLock("n2", "v2", ascertain.MakeLicense("k2", "t2"))
	dl3 := ascertain.MakeDependencyLock("n3", "v3", ascertain.MakeLicense("k3", "t3"))

	dlSlice := []ascertain.DependencyLock{dl1, dl2, dl3}
	dlMap := ascertain.DepLocksToDepLockMap(dlSlice)

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
		result, err := ascertain.MarshalDependencyLocks(DependencyLocks)

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
		result, err := ascertain.UnmarshalDependencyLocks(bytes)

		assert.Nil(t, err)
		assert.Equal(t, DependencyLocks, result)
	})
	t.Run("error on unmarshaling", func(t *testing.T) {
		bytes := []byte("{\"bad\":\"json\"}")
		_, err := ascertain.UnmarshalDependencyLocks(bytes)

		assert.EqualError(t, err,
			"json: cannot unmarshal object into "+
				"Go value of type []ascertain.DependencyLock")
	})
}

func TestLocksSerialization(t *testing.T) {
	reader := reader.NewLimitedReader()

	bytes, err := reader.ReadFileFromPath("./testdata/lock/deserialized.json")
	assert.Nil(t, err)
	expected, err := ascertain.UnmarshalDependencyLocks(bytes)
	assert.Nil(t, err)

	bytes, err = reader.ReadFileFromPath("./testdata/lock/serialized.txt")
	assert.Nil(t, err)
	result, err := ascertain.DeserializeLocks(bytes)
	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}
