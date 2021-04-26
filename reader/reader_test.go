package reader_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/blocky/adlr/reader"
)

const (
	LimitedReaderHappyPath       = "./testdata/happypath.json"
	LimitedReaderMissingFilePath = "./testdata/unicorn"
	LimitedReaderExceedsSizePath = "./testdata/exceedssize.json"

	HappyPathJSON = "{\n   \"fruit\":\"apple\",\n   \"color\":\"red\",\n   \"quantity\":4\n}"
)

func TestLimitedReaderReadFile(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		f, _ := os.Open(LimitedReaderHappyPath)
		r := reader.NewLimitedReaderFromRaw(reader.Kilobyte)
		bytes, err := r.ReadFile(f)

		assert.Nil(t, err)
		assert.Equal(t, HappyPathJSON, string(bytes))
	})
	t.Run("error on nil file", func(t *testing.T) {
		r := reader.NewLimitedReader()
		_, err := r.ReadFile(nil)

		assert.EqualError(t, err, "file pointer is nil")
	})
}

func TestLimitedReaderReadFileFromPath(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		r := reader.NewLimitedReaderFromRaw(reader.Kilobyte)
		bytes, err := r.ReadFileFromPath(LimitedReaderHappyPath)

		assert.Nil(t, err)
		assert.Equal(t, HappyPathJSON, string(bytes))
	})
	t.Run("error on missing file", func(t *testing.T) {
		r := reader.NewLimitedReaderFromRaw(reader.Kilobyte)
		_, err := r.ReadFileFromPath(LimitedReaderMissingFilePath)

		assert.EqualError(t, err, "open ./testdata/unicorn: no such file or directory")
	})
	t.Run("error on exceeds file size", func(t *testing.T) {
		r := reader.NewLimitedReaderFromRaw(reader.Kilobyte)
		_, err := r.ReadFileFromPath(LimitedReaderExceedsSizePath)

		assert.EqualError(t, err, "Refusing to open file of size: 2835")
	})
}
