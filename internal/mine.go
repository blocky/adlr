package internal

import "github.com/go-enry/go-license-detector/v4/licensedb"

// Mines hold the data required to attempt to automatically
// determine a golang module's license
type Mine struct {
	Name    string
	Dir     string
	Version string
	ErrStr  string
	Matches []licensedb.Match
}

// Create a Mine
func MakeMine(
	name string,
	dir string,
	version string,
	matches []licensedb.Match,
) Mine {
	return Mine{
		Name:    name,
		Dir:     dir,
		Version: version,
		Matches: matches,
	}
}

// Add an error string
func (m *Mine) AddError(
	err error,
) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	m.ErrStr = errStr
}
