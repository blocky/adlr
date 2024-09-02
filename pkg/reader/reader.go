package reader

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Constants for data sizes in Bytes
const (
	Byte     int64 = 1
	Kilobyte       = 1000 * Byte
)

// An all-or-nothing reader that attempts to read into memory
// an entire file with a explicit filesize safety catch
type LimitedReader struct {
	maxByteSize int64
}

// Create a LimitedReader with a default maximum filesize of 32 Kilobytes
func NewLimitedReader() *LimitedReader {
	return NewLimitedReaderFromRaw(
		32 * Kilobyte,
	)
}

// Create a LimitedReader with a chosen maximum filesize
func NewLimitedReaderFromRaw(
	maxByteSize int64,
) *LimitedReader {
	return &LimitedReader{maxByteSize}
}

// Read and return file contents from a relative or absolute path
func (l LimitedReader) ReadFileFromPath(
	filepath string,
) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return l.ReadFile(file)
}

// Read and return file contents from a provided file
func (l LimitedReader) ReadFile(
	file *os.File,
) ([]byte, error) {
	err := checkSize(file, l.maxByteSize)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func checkSize(
	file *os.File,
	maxByteSize int64,
) error {
	if file == nil {
		return errors.New("file pointer is nil")
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	if fi.Size() > maxByteSize {
		return fmt.Errorf(
			"Refusing to open file of size: %d", fi.Size())
	}
	return nil
}
