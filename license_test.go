package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

func TestMakeLicense(t *testing.T) {
	l := adlr.MakeLicense("MIT", "license text")
	assert.Equal(t, "MIT", l.Kind)
	assert.Equal(t, "license text", l.Text)
}