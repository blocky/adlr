package gotool

import (
	"encoding/json"
	"errors"
	"io"
)

// Decodes a json build list of modules produced by:
// $ go list -m -json all
type BuildListParser struct{}

// Create BuildListParser with that accepts a io.Reader interface
// and decodes a json list of modules with a json.Decoder
func MakeBuildListParser() BuildListParser {
	return BuildListParser{}
}

func (bl *BuildListParser) newDecoder(
	reader io.Reader,
) (*json.Decoder, error) {
	if reader == nil {
		return nil, errors.New("given nil reader")
	}
	return json.NewDecoder(reader), nil
}

// Parse the json list of modules and return
// a list of unmarshaled Module structs
func (bl *BuildListParser) ParseModuleList(
	reader io.Reader,
) ([]Module, error) {
	var modules []Module

	decoder, err := bl.newDecoder(reader)
	if err != nil {
		return modules, err
	}
	for {
		var m Module
		err := decoder.Decode(&m)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		modules = append(modules, m)
	}
	return modules, nil
}
