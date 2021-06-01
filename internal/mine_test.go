package internal_test

import (
	"errors"
	"testing"

	"github.com/go-enry/go-license-detector/v4/licensedb"
	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/internal"
)

func TestMakeMine(t *testing.T) {
	matches := []licensedb.Match{
		licensedb.Match{License: "MIT", File: "license"},
		licensedb.Match{License: "Apache2,0", File: "License"},
	}
	m := internal.MakeMine("name", "dir", "version", matches)

	assert.Equal(t, "name", m.Name)
	assert.Equal(t, "dir", m.Dir)
	assert.Equal(t, "version", m.Version)
	assert.Equal(t, matches, m.Matches)
}

func TestMineAddError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		var m internal.Mine
		err := errors.New("error")
		m.AddError(err)

		assert.Equal(t, "error", m.ErrStr)
	})
	t.Run("nil error", func(t *testing.T) {
		var m internal.Mine
		var err error
		m.AddError(err)

		assert.Equal(t, "", m.ErrStr)
	})
}
