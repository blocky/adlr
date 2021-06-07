package api

import "github.com/blocky/adlr/pkg/ascertain"

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

// Whitelist is a type alias for ascertain.Whitelist
type Whitelist = ascertain.Whitelist

// MakeWhitelist creates a Whitelist from specified licenses
func MakeWhitelist(
	licenses []string,
) Whitelist {
	return ascertain.MakeLicenseWhitelist(licenses)
}
