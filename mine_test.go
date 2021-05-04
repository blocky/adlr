package adlr_test

import (
	"errors"
	"testing"

	"github.com/go-enry/go-license-detector/v4/licensedb"
	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

func TestMakeMine(t *testing.T) {
	matches := []licensedb.Match{}
	m := adlr.MakeMine("path", "version", matches)

	assert.Equal(t, "path", m.Path)
	assert.Equal(t, "version", m.Version)
	assert.Equal(t, matches, m.Matches)
}

func TestMineAddError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		var m adlr.Mine
		err := errors.New("error")
		m.AddError(err)

		assert.Equal(t, "error", m.ErrStr)
	})
	t.Run("nil error", func(t *testing.T) {
		var m adlr.Mine
		var err error
		m.AddError(err)

		assert.Equal(t, "", m.ErrStr)
	})
}
