package ascertain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/pkg/ascertain"
)

func TestMakeLicense(t *testing.T) {
	l := ascertain.MakeLicense("MIT", "license text")
	assert.Equal(t, "MIT", l.Kind)
	assert.Equal(t, "license text", l.Text)
}
