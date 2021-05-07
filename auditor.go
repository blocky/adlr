package adlr

import "errors"

// Message preface for a non-whitelisted license
const NonWhitelistedLicenseErr = "non-whitelisted license: "

// LicenseAuditor audits DependencyLocks for
// license types not included in the whitelist
type LicenseAuditor struct {
	whitelist LicenseWhitelist
}

// Create a LicenseAuditor with default values
func MakeLicenseAuditor() LicenseAuditor {
	whitelist := MakeLicenseWhitelist()
	return MakeLicenseAuditorFromRaw(whitelist)
}

// Create a LicenseAuditor from specified values
func MakeLicenseAuditorFromRaw(
	whitelist LicenseWhitelist,
) LicenseAuditor {
	return LicenseAuditor{whitelist}
}

// Audit a list of DependencyLocks for license types not included
// in the LicenseWhitelist. Return an error including a list of
// non-whitelist-license DependencyLocks for error printout
func (auditor LicenseAuditor) Audit(
	locks []DependencyLock,
) error {
	var auditErrs []DependencyLock

	for _, lock := range locks {
		err := auditor.AuditLock(lock)
		if err != nil {
			lock.AddErrStr(err.Error())
			auditErrs = append(auditErrs, lock)
		}
	}
	if len(auditErrs) != 0 {
		return &LicenseAuditError{auditErrs}
	}
	return nil
}

// Audit a DependencyLock's license type against the LicenseWhitelist
func (auditor LicenseAuditor) AuditLock(
	lock DependencyLock,
) error {
	var license = lock.License.Kind

	if !auditor.whitelist.Find(license) {
		return errors.New(NonWhitelistedLicenseErr + license)
	}
	return nil
}
