package gotool

import "encoding/json"

// This Module struct has been modified from:
// https://golang.org/src/cmd/go/internal/list/list.go
// This allows modules listed in the buildlist via command:
// $ go list -m -json all
// to be unmarshaled
type Module struct {
	Path     string
	Version  string
	Replace  *Module
	Main     bool
	Indirect bool
	Dir      string
}

// Unmarshal a buildlist module from json
func (m *Module) UnmarshalJSON(bytes []byte) error {
	type Alias Module
	tmp := struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	return json.Unmarshal(bytes, &tmp)
}

func FilterImportModules(
	modules []Module,
) []Module {
	var direct []Module
	for _, m := range modules {
		if m.Main == true {
			continue
		}
		direct = append(direct, m)
	}
	return direct
}
