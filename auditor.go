package adlr

import "errors"

type LicenseAuditor struct {
	whitelist LicenseWhitelist
}

func MakeLicenseAuditor() LicenseAuditor {
	whitelist := MakeLicenseWhitelist()
	return MakeLicenseAuditorFromRaw(whitelist)
}

func MakeLicenseAuditorFromRaw(
	whitelist LicenseWhitelist,
) LicenseAuditor {
	return LicenseAuditor{whitelist}
}

func (auditor LicenseAuditor) AuditLocks(
	locks []DependencyLock,
) error {
	var auditErrs []error
	for _, lock := range locks {
		err := auditor.AuditLock(lock)
		if err != nil {
			auditErrs = append(auditErrs, err)
		}
	}
	if len(auditErrs) != 0 {
		return &LicenseAuditError{auditErrs}
	}
	return nil
}

func (auditor LicenseAuditor) AuditLock(
	lock DependencyLock,
) error {
	var name = lock.Name
	var license = lock.License.Kind

	if !auditor.whitelist.Find(license) {
		return errors.New(name + ": [" + license + "]")
	}
	return nil
}
