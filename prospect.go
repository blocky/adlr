package adlr

import "github.com/blocky/adlr/gotool"

type Prospect struct {
	Path    string
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
			mod.Version,
		)
	}
	return prospects
}

func MakeProspect(
	path, version string,
) Prospect {
	return Prospect{
		Path:    path,
		Version: version,
	}
}

func (p *Prospect) AddError(
	err error,
) {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	p.ErrStr = errStr
}
