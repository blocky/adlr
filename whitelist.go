package adlr

import "sort"

// The following list follows SPDX License Identifier standards
// To add an ACCEPTABLE LICENSE to this list, see:
// https://spdx.org/licenses/
// for the license identifier

var Whitelist = []string{
	"Apache-2.0",
	"BSD-1-Clause",
	"BSD-2-Clause",
	"BSD-3-Clause",
	"MIT",
	"MIT-0",
}

type LicenseWhitelist struct {
	whitelist []string
	init      bool
}

func MakeLicenseWhitelist() LicenseWhitelist {
	return MakeLicenseWhitelistFromRaw(Whitelist)
}

func MakeLicenseWhitelistFromRaw(
	whitelist []string,
) LicenseWhitelist {
	init := preprocess(whitelist)
	return LicenseWhitelist{whitelist, init}
}

func (lw LicenseWhitelist) Find(
	license string,
) bool {
	return lw.init && find(lw.whitelist, license)
}

func find(
	list []string,
	x string,
) bool {
	i := sort.SearchStrings(list, x)
	if i == len(list) { // if after every list item
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
