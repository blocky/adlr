package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

func TestMakeDependencyLock(t *testing.T) {
	l := adlr.MakeLicense("kind", "text")
	dl := adlr.MakeDependencyLock("name", "version", l)

	assert.Equal(t, l, dl.License)
	assert.Equal(t, "name", dl.Name)
	assert.Equal(t, "version", dl.Version)
}
