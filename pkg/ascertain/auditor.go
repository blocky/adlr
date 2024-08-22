package ascertain

import "errors"

// Message preface for a non-whitelisted license
const NonWhitelistedLicenseErr = "non-whitelisted license: "

// LicenseAuditor audits DependencyLocks for
// license types not included in the whitelist
type LicenseAuditor struct {
	whitelist Whitelist
}

// Create a LicenseAuditor from specified Whitelist
func MakeLicenseAuditor(
	whitelist Whitelist,
) LicenseAuditor {
	return LicenseAuditor{whitelist}
}

// Audit a list of DependencyLocks for license types not included
// in the LicenseWhitelist. Return an error including a list of
// non-whitelist-license DependencyLocks for error printout
func (auditor LicenseAuditor) Audit(
	locks ...DependencyLock,
) ([]DependencyLock, error) {
	var auditErrs []DependencyLock

	var verified = make([]DependencyLock, 0)
	for _, lock := range locks {
		err := auditor.AuditLock(lock)
		if err != nil {
			lock.AddErrStr(err.Error())
			auditErrs = append(auditErrs, lock)
			continue
		}
		verified = append(verified, lock)
	}
	if len(auditErrs) != 0 {
		return verified, &LicenseAuditError{auditErrs}
	}
	return verified, nil
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
