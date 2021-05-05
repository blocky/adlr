package adlr

import "github.com/blocky/adlr/gotool"

type Prospect struct {
	Name    string
	Dir     string
	Version string
	ErrStr  string
}

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

func (p *Prospect) AddErrStr(
	errStr string,
) {
	p.ErrStr = errStr
}
