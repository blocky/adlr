package ascertain

import "github.com/blocky/adlr/pkg/gotool"

// Prospects hold the data required to text mine a golang module
type Prospect struct {
	Name    string
	Dir     string
	Version string
	ErrStr  string
}

// Create a list of Prospect from a list of gotool.Module,
// carrying-over all important data
func MakeProspects(
	modules ...gotool.Module,
) []Prospect {
	var prospects = make([]Prospect, len(modules))
	for i, mod := range modules {
		prospects[i] = MakeProspect(
			mod.Path,
			mod.Dir,
			mod.Version,
		)
	}
	return prospects
}

// Create a Prospect
func MakeProspect(
	name string,
	dir string,
	version string,
) Prospect {
	return Prospect{
		Name:    name,
		Dir:     dir,
		Version: version,
	}
}

// Add error string
func (p *Prospect) AddErrStr(
	errStr string,
) {
	p.ErrStr = errStr
}
