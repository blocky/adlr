package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

func TestProspect(t *testing.T) {
	p := adlr.MakeProspect("name", "dir", "version")

	assert.Equal(t, "name", p.Name)
	assert.Equal(t, "dir", p.Dir)
	assert.Equal(t, "version", p.Version)
}

func TestProspectAddErrStr(t *testing.T) {
	var p adlr.Prospect
	p.AddErrStr("error")

	assert.Equal(t, "error", p.ErrStr)
}
