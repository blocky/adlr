package internal

import "sort"

// This list is of SPDX License Identifiers, the standard
// used by the text-mining package:
// github.com/go-enry/go-license-detector/v4
// and contains licenses deemed automatically fulfillable.
// To add to this list, see:
// https://spdx.org/licenses/
// for license identifiers
var DefaultWhitelist = []string{
	"Apache-2.0",
	"BSD-1-Clause",
	"BSD-2-Clause",
	"BSD-3-Clause",
	"MIT",
	"MIT-0",
}

type Whitelist interface {
	Find(string) bool
}

// LicenseWhitelist is a whitelist of automatically fulfillable
// licenses, with search by license type
type LicenseWhitelist struct {
	whitelist []string
	init      bool
}

// Create a LicenseWhitelist with default values
func MakeLicenseWhitelist() Whitelist {
	return MakeLicenseWhitelistFromRaw(DefaultWhitelist)
}

// Create a LicenseWhitelist from specified values.
// Initialize whitelist for searching
func MakeLicenseWhitelistFromRaw(
	whitelist []string,
) Whitelist {
	init := preprocess(whitelist)
	return LicenseWhitelist{whitelist, init}
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
