package api

import "github.com/blocky/adlr/internal"

// Whitelist is a type alias for internal.Whitelist
type Whitelist = internal.Whitelist

// MakeWhitelist creates a default Whitelist
func MakeWhitelist() Whitelist {
	return internal.MakeLicenseWhitelist()
}

// MakeWhitelistFromRaw creates a Whitelist from specified licenses
func MakeWhitelistFromRaw(
	licenses []string,
) Whitelist {
	return internal.MakeLicenseWhitelistFromRaw(licenses)
}
