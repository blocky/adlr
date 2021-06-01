package api

import "github.com/blocky/adlr/internal"

type Prospector interface {
	Prospect(...Prospect) ([]Mine, error)
}

// Create a LicenseProspector
func MakeProspector() Prospector {
	return internal.MakeLicenseProspector()
}
