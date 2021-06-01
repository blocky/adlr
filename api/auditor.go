package api

import "github.com/blocky/adlr/internal"

// Auditor takes a variadic list of DependencyLocks and audits their licenses
// against a license whitelist, returning an error of all offending
// DependencyLocks and their non-whitelisted licenses
type Auditor interface {
	Audit(...DependencyLock) error
}

// MakeAuditor creates a default Auditor with a default Whitelist
func MakeAuditor(
	whitelist Whitelist,
) Auditor {
	return internal.MakeLicenseAuditor(whitelist)
}
