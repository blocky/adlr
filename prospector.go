package adlr

import "github.com/go-enry/go-license-detector/v4/licensedb"

type LicenseProspector struct{}

func MakeLicenseProspector() LicenseProspector {
	return LicenseProspector{}
}

func (lp LicenseProspector) Prospect(
	deps ...Dependency,
) ([]Dependency, error) {
	var prospectErrs []licensedb.Result

	for i, dep := range deps {
		result := lp.ProspectLicense(dep.Module.Dir)
		deps[i].AddResult(result)

		if result.ErrStr != "" { // could not find dir or license files
			prospectErrs = append(prospectErrs, result)
		}
	}
	if len(prospectErrs) != 0 {
		return deps, &LicenseProspectingError{prospectErrs}
	}
	return deps, nil
}

func (lp LicenseProspector) ProspectLicense(
	dir string,
) licensedb.Result {
	// always returns n=1 array on one arg
	return licensedb.Analyse(dir)[0]
}
