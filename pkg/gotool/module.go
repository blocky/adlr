package gotool

import (
	"encoding/json"

	"github.com/samber/lo"
)

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
	return lo.Filter(modules, func(module Module, _ int) bool {
		return !module.Main
	})
}

func RemoveExemptModules(
	modules []Module,
	exempt []string,
) []Module {
	return lo.Filter(modules, func(module Module, _ int) bool {
		for _, e := range exempt {
			if module.Path == e {
				return false
			}
		}
		return true
	})
}
