package api

import "github.com/blocky/adlr/internal"

type Whitelist = internal.Whitelist

func MakeWhitelist() Whitelist {
	return internal.MakeLicenseWhitelist()
}

func MakeWhitelistFromRaw(
	licenses []string,
) Whitelist {
	return internal.MakeLicenseWhitelistFromRaw(licenses)
}
