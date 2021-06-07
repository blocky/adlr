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

// Filter out main, replaced, and indirect modules from
// a buildlist. Return a new list of directly imported modules
func FilterDirectImportModules(
	modules []Module,
) []Module {
	var direct []Module
	for _, m := range modules {
		if m.Main == true {
			continue
		} else if m.Replace != nil {
			// Any modules replaced in a go.mod with
			// the keyword "Replace" will be caught here
			continue
		} else if m.Indirect == false {
			direct = append(direct, m)
		}
	}
	return direct
}
