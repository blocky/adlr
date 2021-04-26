package gotool_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/gotool"
)

var TestDirBuildList = "./testdata/buildlist/"

var ModuleList = []gotool.Module{
	gotool.Module{
		"cloud.google.com/go",
		"v0.33.1", nil, false, true,
		"",
	},
	gotool.Module{
		"github.com/BurntSushi/toml",
		"v0.3.1", nil, false, true,
		"/home/user/go/pkg/mod/github.com/!burnt!sushi/toml@v0.3.1",
	},
	gotool.Module{
		"github.com/DATA-DOG/go-sqlmock",
		"v1.5.0", nil, false, true,
		"/home/user/go/pkg/mod/github.com/!d!a!t!a-!d!o!g/go-sqlmock@v1.5.0",
	},
}

func TestBuildListParser(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		f, _ := os.Open(TestDirBuildList + "buildlist.json")
		p := gotool.MakeBuildListParser()
		list, err := p.ParseModuleList(f)

		assert.Nil(t, err)
		assert.Equal(t, ModuleList, list)
	})
	t.Run("error on nil file", func(t *testing.T) {
		p := gotool.MakeBuildListParser()
		_, err := p.ParseModuleList(nil)

		assert.EqualError(t, err, "given nil reader")
	})
	t.Run("error on bad json", func(t *testing.T) {
		r := strings.NewReader(`bad json`)
		p := gotool.MakeBuildListParser()
		_, err := p.ParseModuleList(r)

		assert.EqualError(t, err,
			"invalid character 'b' looking for beginning of value")
	})
}
