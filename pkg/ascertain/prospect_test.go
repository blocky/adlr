package ascertain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
)

func TestProspect(t *testing.T) {
	p := ascertain.MakeProspect("name", "dir", "version")

	assert.Equal(t, "name", p.Name)
	assert.Equal(t, "dir", p.Dir)
	assert.Equal(t, "version", p.Version)
}

func TestProspectAddErrStr(t *testing.T) {
	var p ascertain.Prospect
	p.AddErrStr("error")

	assert.Equal(t, "error", p.ErrStr)
}
