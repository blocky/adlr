package api

import "github.com/blocky/adlr/pkg/ascertain"

// Auditor takes a variadic list of DependencyLocks and audits their licenses
// against a license whitelist, returning an error of all offending
// DependencyLocks and their non-whitelisted licenses
type Auditor interface {
	Audit(...DependencyLock) error
}

// MakeAuditor creates an Auditor with a specified Whitelist
func MakeAuditor(
	whitelist Whitelist,
) Auditor {
	return ascertain.MakeLicenseAuditor(whitelist)
}
