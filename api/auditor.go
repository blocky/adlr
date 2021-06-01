package api

import "github.com/blocky/adlr/internal"

type Auditor interface {
	Audit([]DependencyLock) error
}

func MakeAuditor() Auditor {
	return internal.MakeLicenseAuditor()
}

func MakeAuditorFromRaw(
	whitelist Whitelist,
) Auditor {
	return internal.MakeLicenseAuditorFromRaw(whitelist)
}
