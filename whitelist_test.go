package adlr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr"
)

var fruitlist = []string{
	"apricot",
	"blueberry",
	"cherry",
}

func TestLicenseWhitelist(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(fruitlist)
		assert.True(t, w.Find("apricot"))
		assert.True(t, w.Find("blueberry"))
		assert.True(t, w.Find("cherry"))
	})
	t.Run("false on not exist", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw(fruitlist)
		assert.False(t, w.Find("apple"), "before first item")
		assert.False(t, w.Find("avocado"), "after first item")
		assert.False(t, w.Find("banana"), "before second item")
		assert.False(t, w.Find("boysenberry"), "after second item")
		assert.False(t, w.Find("cantalope"), "before third item")
		assert.False(t, w.Find("cucumber"), "after third item")
	})
	t.Run("false on empty list", func(t *testing.T) {
		w := adlr.MakeLicenseWhitelistFromRaw([]string{})
		assert.False(t, w.Find("unicorn"))
	})
}
