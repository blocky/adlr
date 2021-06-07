package api

import "github.com/blocky/adlr/pkg/ascertain"

// Prospector takes a variadic list of Prospects and uses text mining to
// derive Mines with potential matches containing license type, license
// file name, and confidence float, and returns an error of failed Mines
// from missing directories or missing license files
type Prospector interface {
	Prospect(...Prospect) ([]Mine, error)
}

// MakeProspector creates a Prospector
func MakeProspector() Prospector {
	return ascertain.MakeLicenseProspector()
}
