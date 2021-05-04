package adlr

import "github.com/go-enry/go-license-detector/v4/licensedb"

type Mine struct {
	Path    string
	Version string
	ErrStr  string
	Matches []licensedb.Match
}

func MakeMine(
	path string,
	version string,
	matches []licensedb.Match,
) Mine {
	return Mine{
		Path:    path,
		Version: version,
		Matches: matches,
	}
}

func (m *Mine) AddError(
	err error,
) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	m.ErrStr = errStr
}
