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
	{
		Path:    "bky.sh/pallets/nub",
		Version: "v0.0.0",
		Replace: &gotool.Module{
			Path: "../nub",
			Dir:  "/home/user/Work/bky.sh/pallets/nub",
		},
		Main:     false,
		Indirect: false,
		Dir:      "/home/user/Work/bky.sh/pallets/nub",
	},
	{
		Path:     "cirello.io/pglock",
		Version:  "v1.8.1-0.20200922151210-b76be34db4ac",
		Main:     false,
		Indirect: true,
		Dir:      "/home/user/go/pkg/mod/cirello.io/pglock@v1.8.1-0.20200922151210-b76be34db4ac",
	},
	{
		Path:     "github.com/spf13/cobra",
		Version:  "v0.0.5",
		Main:     false,
		Indirect: false,
		Dir:      "/home/user/go/pkg/mod/github.com/spf13/cobra@v0.0.5",
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

func TestFilterImportModules(t *testing.T) {
	t.Run("happy path - filter all but main", func(t *testing.T) {
		f, _ := os.Open(TestDirModule + "all_imports.json")
		p := gotool.MakeBuildListParser()
		list, err := p.ParseModuleList(f)
		assert.Nil(t, err)

		result := gotool.FilterImportModules(list)
		assert.Equal(t, Modules, result)
	})
	t.Run("happy path - empty on only main import", func(t *testing.T) {
		f, _ := os.Open(TestDirModule + "main_import.json")
		p := gotool.MakeBuildListParser()
		list, err := p.ParseModuleList(f)
		assert.Nil(t, err)

		result := gotool.FilterImportModules(list)
		assert.Equal(t, []gotool.Module(nil), result)
	})
}
