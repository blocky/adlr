package internal

import "sort"

type Whitelist interface {
	Find(string) bool
}

// LicenseWhitelist is a whitelist of automatically fulfillable
// licenses, with search by license type
type LicenseWhitelist struct {
	whitelist []string
	init      bool
}

// Create a LicenseWhitelist with a list of licenses
func MakeLicenseWhitelist(
	licenses []string,
) Whitelist {
	init := preprocess(licenses)
	return LicenseWhitelist{licenses, init}
}

// Search a whitelist of license types that a license exists
func (lw LicenseWhitelist) Find(
	license string,
) bool {
	return lw.init && find(lw.whitelist, license)
}

func find(
	list []string,
	x string,
) bool {
	// return index i to insert item x
	// x exists if list[i] == x
	i := sort.SearchStrings(list, x)
	if i >= len(list) {
		return false
	} else if x != list[i] {
		return false
	}
	return true
}

func preprocess(list []string) bool {
	sort.Strings(list)
	return sort.StringsAreSorted(list)
}
