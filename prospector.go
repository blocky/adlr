package adlr

import "github.com/go-enry/go-license-detector/v4/licensedb"

// Text mine Prospects and return licensedb.Results
type LicenseProspector struct{}

// Create a LicenseProspector
func MakeLicenseProspector() LicenseProspector {
	return LicenseProspector{}
}

// Concurrently text mine a list of Prospect, and generate a
// a stable list of licensedb.Result. Return a stable list of Mine
// with its respective list of licensedb.Match from text mining.
// An error return occurs for missing directories or missing licenses,
// and should be treated as non-recoverable
func (lp LicenseProspector) Prospect(
	prospects ...Prospect,
) ([]Mine, error) {
	var prospectErrs []Prospect

	var paths = make([]string, len(prospects))
	for i, prospect := range prospects {
		paths[i] = prospect.Dir
	}
	// compute results concurrently; stable algorithm
	results := lp.ProspectLicenses(paths...)

	var mined = make([]Mine, len(prospects))
	for i, prospect := range prospects {
		var result = results[i]

		if result.ErrStr != "" { // could not find dir or license files
			prospect.AddErrStr(result.ErrStr)
			prospectErrs = append(prospectErrs, prospect)
		}
		mined[i] = MakeMine(
			prospect.Name,
			prospect.Dir,
			prospect.Version,
			result.Matches,
		)
	}
	if len(prospectErrs) != 0 {
		return mined, &LicenseProspectingError{prospectErrs}
	}
	return mined, nil
}

// Concurrently text mine paths and return a stable list of licensedb.Result
func (lp LicenseProspector) ProspectLicenses(
	paths ...string,
) []licensedb.Result {
	return licensedb.Analyse(paths...)
}
