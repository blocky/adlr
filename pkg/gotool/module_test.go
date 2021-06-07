package gotool_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/gotool"
)

const TestDirModule = "./testdata/module/"

var ModuleStr = `{` +
	`"Path":"path/",` +
	`"Version":"12345",` +
	`"Replace":null,` +
	`"Main":true,` +
	`"Indirect":true,` +
	`"Dir":"/home/user/path"}`

var Modules = []gotool.Module{
	gotool.Module{
		Path:    "github.com/spf13/cobra",
		Version: "v0.0.5",
		Dir:     "/home/user/go/pkg/mod/github.com/spf13/cobra@v0.0.5",
	},
}

func TestModuleUnmarshal(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := gotool.Module{
			Path:     "path/",
			Version:  "12345",
			Replace:  nil,
			Main:     true,
			Indirect: true,
			Dir:      "/home/user/path",
		}
		var m gotool.Module
		err := m.UnmarshalJSON([]byte(ModuleStr))

		assert.Nil(t, err)
		assert.Equal(t, expected, m)
	})
	t.Run("error on bad json", func(t *testing.T) {
		var m gotool.Module
		err := m.UnmarshalJSON([]byte("bad json"))

		assert.EqualError(t, err,
			"invalid character 'b' looking for beginning of value")
	})
}

func TestFilterDirectImportModules(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		f, _ := os.Open(TestDirModule + "directimports.json")
		p := gotool.MakeBuildListParser()
		list, err := p.ParseModuleList(f)
		assert.Nil(t, err)

		result := gotool.FilterDirectImportModules(list)
		assert.Equal(t, Modules, result)
	})
	t.Run("empty on no direct imports", func(t *testing.T) {
		f, _ := os.Open(TestDirModule + "indirectimports.json")
		p := gotool.MakeBuildListParser()
		list, err := p.ParseModuleList(f)
		assert.Nil(t, err)

		result := gotool.FilterDirectImportModules(list)
		assert.Equal(t, []gotool.Module(nil), result)
	})
}
