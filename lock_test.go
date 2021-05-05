package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
	"github.com/blocky/adlr/reader"
)

var DependencyLocks = []adlr.DependencyLock{
	adlr.DependencyLock{
		Name:    "github.com/spf13/viper",
		Version: "v1.4.0",
		License: adlr.License{
			"MIT",
			"MIT License",
		},
	},
	adlr.DependencyLock{
		Name:    "github.com/stretchr/testify",
		Version: "v1.6.1",
		License: adlr.License{
			"MIT",
			"MIT License",
		},
	},
	adlr.DependencyLock{
		Name:    "github.com/ybbus/jsonrpc",
		Version: "v2.1.2+incompatible",
		License: adlr.License{
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

func TestDependencyLockAddErrStr(t *testing.T) {
	var l adlr.DependencyLock
	l.AddErrStr("error")

	assert.Equal(t, "error", l.ErrStr)
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

func TestMarshalDependencyLocks(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		bytes, err := reader.
			NewLimitedReader().
			ReadFileFromPath("./testdata/lock/marshaled.json")
		assert.Nil(t, err)
		result, err := adlr.MarshalDependencyLocks(DependencyLocks)

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
		result, err := adlr.UnmarshalDependencyLocks(bytes)

		assert.Nil(t, err)
		assert.Equal(t, DependencyLocks, result)
	})
	t.Run("error on unmarshaling", func(t *testing.T) {
		bytes := []byte("{\"bad\":\"json\"}")
		_, err := adlr.UnmarshalDependencyLocks(bytes)

		assert.EqualError(t, err,
			"json: cannot unmarshal object into "+
				"Go value of type []adlr.DependencyLock")
	})
}

func TestLocksSerialization(t *testing.T) {
	reader := reader.NewLimitedReader()

	bytes, err := reader.ReadFileFromPath("./testdata/lock/deserialized.json")
	assert.Nil(t, err)
	expected, err := adlr.UnmarshalDependencyLocks(bytes)
	assert.Nil(t, err)

	bytes, err = reader.ReadFileFromPath("./testdata/lock/serialized.txt")
	assert.Nil(t, err)
	result, err := adlr.DeserializeLocks(bytes)
	assert.Nil(t, err)

	assert.Equal(t, expected, result)
}
