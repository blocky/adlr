package adlr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

func TestProspect(t *testing.T) {
	p := adlr.MakeProspect("path", "version")

	assert.Equal(t, "path", p.Path)
	assert.Equal(t, "version", p.Version)
}

func TestProspectAddError(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		var p adlr.Prospect
		err := errors.New("error")
		p.AddError(err)

		assert.Equal(t, "error", p.ErrStr)
	})
	t.Run("nil error", func(t *testing.T) {
		var p adlr.Prospect
		var err error
		p.AddError(err)

		assert.Equal(t, "", p.ErrStr)
	})
}
